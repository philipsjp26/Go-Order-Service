// Package appctx
package appctx

import (
    "encoding/json"
    "sync"

    "gitlab.privy.id/order_service/internal/consts"
    "gitlab.privy.id/order_service/pkg/msgx"
)

var (
	rsp    *Response
	oneRsp sync.Once
)

// Response presentation contract object
type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	lang    string      `json:"-"`
	Meta    interface{} `json:"meta,omitempty"`
	msgKey  string
}

// MetaData represent meta data response for multi data
type MetaData struct {
    Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPage  int64 `json:"total_page"`
	TotalCount int64 `json:"total_count"`
}

// GetMessage method to transform response name var to message detail
func (r *Response) GetMessage() string {
	return msgx.Get(r.msgKey, r.lang).Text()
}

// Generate setter message
func (r *Response) Generate() *Response {
	if r.lang == "" {
		r.lang = consts.LangDefault
	}
	msg := msgx.Get(r.msgKey, r.lang)
	if r.Message == nil {
		r.Message = msg.Text()
	}

	if r.Code == 0 {
		r.Code = msg.Status()
	}

	return r
}

// WithCode setter response var name
func (r *Response) WithCode(c int) *Response {
	r.Code = c
	return r
}

// WithData setter data response
func (r *Response) WithData(v interface{}) *Response {
	r.Data = v
	return r
}

// WithError setter error messages
func (r *Response) WithError(v interface{}) *Response {
	r.Errors = v
	return r
}

func (r *Response) WithMsgKey(v string) *Response {
	r.msgKey = v
	return r
}

// WithMeta setter meta data response
func (r *Response) WithMeta(v interface{}) *Response {
	r.Meta = v
	return r
}

// WithLang setter language response
func (r *Response) WithLang(v string) *Response {
	r.lang = v
	return r
}

// WithMessage setter custom message response
func (r *Response) WithMessage(v interface{}) *Response {
	if v != nil {
		r.Message = v
	}

	return r
}

func (r *Response) Byte() []byte {
	if r.Code == 0 || r.Message == nil {
		r.Generate()
	}

	b, _ := json.Marshal(r)
	return b
}

// NewResponse initialize response
func NewResponse() *Response {
	oneRsp.Do(func() {
		rsp = &Response{}
	})

	// clone response
	x := *rsp

	return &x
}
