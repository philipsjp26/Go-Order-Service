package ucase

import (
	"fmt"

	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/internal/consts"
	repository "gitlab.privy.id/order_service/internal/repositories/menu"
	ucase "gitlab.privy.id/order_service/internal/ucase/contract"
	"gitlab.privy.id/order_service/pkg/logger"
)

type Menu struct {
	repo repository.MenuRepository
}

func NewMenu(repo repository.MenuRepository) ucase.UseCase {
	return &Menu{repo: repo}
}

func (r *Menu) Serve(dctx *appctx.Data) appctx.Response {

	res, err := r.repo.ListMenu(dctx.Request.Context())
	if err != nil {
		logger.WarnWithContext(dctx.Request.Context(), fmt.Sprintf("List menu error : %v", err))
		return *appctx.NewResponse().WithMsgKey(consts.RespError)
	}
	return *appctx.NewResponse().WithCode(200).WithMessage("Success").WithData(res)
}
