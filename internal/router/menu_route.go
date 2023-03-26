package router

import (
	"net/http"

	"gitlab.privy.id/order_service/internal/handler"
	repo "gitlab.privy.id/order_service/internal/repositories/menu"
	"gitlab.privy.id/order_service/internal/ucase"
	"gitlab.privy.id/order_service/pkg/databasex"
	"gitlab.privy.id/order_service/pkg/routerkit"
)

func MenuRoute(rtr *router, db *databasex.DB) *routerkit.Router {
	api := rtr.router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	menu := v1.PathPrefix("/menu").Subrouter()

	menuRepository := repo.NewMenuRepository(db)
	menuUseCase := ucase.NewMenu(menuRepository)

	menu.HandleFunc("", rtr.handle(
		handler.HttpRequest,
		menuUseCase,
	)).Methods(http.MethodGet)
	return rtr.router
}
