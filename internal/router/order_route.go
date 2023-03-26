package router

import (
	"net/http"

	"gitlab.privy.id/order_service/internal/handler"
	repo "gitlab.privy.id/order_service/internal/repositories/order"
	"gitlab.privy.id/order_service/internal/ucase"
	"gitlab.privy.id/order_service/pkg/databasex"
	"gitlab.privy.id/order_service/pkg/routerkit"
)

func OrderRoute(rtr *router, db *databasex.DB) *routerkit.Router {
	api := rtr.router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	order := v1.PathPrefix("/order").Subrouter()

	orderRepository := repo.NewOrderRepository(db)
	orderUseCase := ucase.NewOrder(orderRepository)

	order.HandleFunc("", rtr.handle(
		handler.HttpRequest,
		orderUseCase,
	)).Methods(http.MethodPost)
	return rtr.router
}
