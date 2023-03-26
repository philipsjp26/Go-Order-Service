package middleware

import (
	"net/http"

	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/internal/consts"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	rsp := appctx.NewResponse().
		WithMsgKey(consts.RespNoRouteFound).
		Generate()
	w.Header().Set("Content-Type", consts.HeaderContentTypeJSON)
	w.WriteHeader(rsp.Code)
	w.Write(rsp.Byte())
	return
}
