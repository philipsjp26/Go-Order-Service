package builderx

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"gitlab.privy.id/order_service/pkg/util"
)

const (
	placeholder          = "?"
	layoutDateTimeFormat = `2006-01-02 15:04:05`
)

var escapedPlaceholder = strings.Repeat(placeholder, 2)

func isTime(obj reflect.Value) bool {
	_, ok := obj.Interface().(time.Time)
	if ok {
		return ok
	}

	_, ok = obj.Interface().(*time.Time)

	return ok
}

func timeIsZero(obj reflect.Value) bool {
	t, ok := obj.Interface().(time.Time)
	if ok {
		return t.IsZero()
	}

	t2, ok := obj.Interface().(*time.Time)
	if ok {
		return false
	}

	return t2 == nil
}

func isNil(i interface{}) bool {
	if i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil()) {
		return true
	}

	return false
}

func PostgrePlaceholder(n int) string {
	return fmt.Sprintf("$%d", n+1)
}

func MsSqlPlaceholder(n int) string {
	return fmt.Sprintf("@p%d", n+1)
}

func ToPostgrePlaceHolder(query string) string {
	b := strings.Builder{}
	n := 0
	for {
		index := strings.Index(query, placeholder)
		if index == -1 {
			break
		}

		// escape placeholder by repeating it twice
		if strings.HasPrefix(query[index:], escapedPlaceholder) {
			b.WriteString(query[:index]) // Write placeholder once, not twice
			query = strings.TrimSpace(query[index+1:])

			continue
		}

		b.WriteString(query[:index])
		b.WriteString(PostgrePlaceholder(n))
		query = query[index+len(placeholder):]
		n++
	}

	// placeholder not found; write remaining query
	b.WriteString(query)

	return b.String()

}

// StructToKeyValue converts a struct to a key value the struct's tags.
// StructToKeyValue uses tags on struct fields to decide which fields to add to the
// returned slice struct.
func StructToKeyValue(src interface{}, tag string) ([]KeyValue, error) {
	var out []KeyValue
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		if !fi.IsExported() {
			continue
		}

		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {

			if isNil(v.Field(i).Interface()) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && util.IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && tagsv[1] == "omitempty") {
					continue
				}
			}

			if v.Field(i).Kind() == reflect.Struct {
				continue
			}

			col := tagsv[0]

			// set key value of struct key value interface output
			out = append(out, KeyValue{
				Key:   col,
				Value: v.Field(i).Interface(),
			})
		}

		if tagsv[0] == "" && v.Field(i).Kind() == reflect.Struct {
			x, err := StructToKeyValue(v.Field(i).Interface(), tag)
			if err != nil {
				return out, err
			}

			out = append(out, x...)
		}
	}

	return out, nil
}

// StructToMap converts a struct to a map using the struct's tags.
// StructToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func StructToMap(src interface{}, tag string) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		//field := reflectValue.Field(i).Interface()
		if !fi.IsExported() {
			continue
		}

		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {

			if isNil(v.Field(i).Interface()) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && util.IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && tagsv[1] == "omitempty") {
					continue
				}
			}

			if v.Field(i).Kind() == reflect.Struct {
				continue
			}

			col := tagsv[0]

			if util.InArray("ne", tagsv) {
				col = fmt.Sprintf("%s !", col)
			}
			// set key value of map interface output
			out[col] = v.Field(i).Interface()
		}

		if tagsv[0] == "" && v.Field(i).Kind() == reflect.Struct {
			x, err := StructToMap(v.Field(i).Interface(), tag)
			if err != nil {
				return out, err
			}

			for y, z := range x {
				out[y] = z
			}
		}
	}

	return out, nil
}

// ToColumnsValues iterate struct to separate key field and value
func ToColumnsValues(src interface{}, tag string) ([]string, []interface{}, error) {
	var columns []string
	var values []interface{}

	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {

			if tagsv[0] == "-" {
				continue
			}

			if isNil(v.Field(i).Interface()) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && util.InArray("omitempty", tagsv)) && util.IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && util.InArray("omitempty", tagsv)) {
					continue
				}
			}

			col := tagsv[0]

			if util.InArray("ne", tagsv) {
				col = fmt.Sprintf("%s !", col)
			}

			// set value of string slice to value in struct field
			columns = append(columns, col)

			// set value interface of value struct field
			values = append(values, v.Field(i).Interface())

		}
	}

	return columns, values, nil
}

// StructToKeyValueWithSkipOmitEmpty converts a struct to a key value the struct's tags.
// StructToKeyValueWithSkipOmitEmpty uses tags on struct fields to decide which fields to add to the
// returned slice struct.
func StructToKeyValueWithSkipOmitEmpty(src interface{}, tag string, columns []string, skipOmitEmpty bool) ([]KeyValue, error) {
	var out []KeyValue
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		//field := reflectValue.Field(i).Interface()
		if !fi.IsExported() {
			continue
		}

		tagsv := strings.Split(fi.Tag.Get(tag), ",")
		col := tagsv[0]
		if col != "" && fi.PkgPath == "" {

			if isNil(v.Field(i).Interface()) {
				continue
			}

			if !util.InArray(col, columns) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && util.IsEmptyValue(v.Field(i).Interface()) && skipOmitEmpty {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && tagsv[1] == "omitempty") && skipOmitEmpty {
					continue
				}
			}

			if v.Field(i).Kind() == reflect.Struct {
				continue
			}

			// set key value of struct key value interface output
			out = append(out, KeyValue{
				Key:   col,
				Value: v.Field(i).Interface(),
			})
		}

		if col == "" && v.Field(i).Kind() == reflect.Struct {
			x, err := StructToKeyValue(v.Field(i).Interface(), tag)
			if err != nil {
				return out, err
			}

			out = append(out, x...)
		}
	}

	return out, nil
}
