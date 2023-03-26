package ucase

import (
	"gitlab.privy.id/order_service/internal/appctx"
	repostory "gitlab.privy.id/order_service/internal/repositories/order"
	ucase "gitlab.privy.id/order_service/internal/ucase/contract"
)

type Order struct {
	repo repostory.OrderRepository
}

func NewOrder(repo repostory.OrderRepository) ucase.UseCase {
	return &Order{repo: repo}
}

func (repo *Order) Serve(dctx *appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(200)
}
