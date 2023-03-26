package builderx

type (
	KeyValue struct {
		Key   string
		Value interface{}
	}

	QueryWhere struct {
		Columns []string
		Values  []interface{}
		Limit   int64
		Page    int64
		Query   string
	}
)
