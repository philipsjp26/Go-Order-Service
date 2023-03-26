package ucase

import (
	"fmt"

	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/internal/consts"
	"gitlab.privy.id/order_service/internal/entity"
	repostory "gitlab.privy.id/order_service/internal/repositories/order"
	ucase "gitlab.privy.id/order_service/internal/ucase/contract"
	"gitlab.privy.id/order_service/pkg/logger"
)

type Order struct {
	repo repostory.OrderRepository
}

func NewOrder(repo repostory.OrderRepository) ucase.UseCase {
	return &Order{repo: repo}
}

func (r *Order) Serve(dctx *appctx.Data) appctx.Response {
	var (
		param entity.OrderRequest
	)

	err := dctx.Cast(&param)

	if err != nil {
		return *appctx.NewResponse().WithMsgKey(consts.RespError)
	}

	res, err := r.repo.Order(dctx.Request.Context(), param)
	if err != nil {
		logger.WarnWithContext(dctx.Request.Context(), fmt.Sprintf("store error: %v", err))
		return *appctx.NewResponse().WithMsgKey(consts.RespError)
	}

	return *appctx.NewResponse().WithCode(201).WithMessage("Created").WithData(res)
}
