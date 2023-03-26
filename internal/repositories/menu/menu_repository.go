package menu

import (
	"context"

	"gitlab.privy.id/order_service/internal/entity"
)

type MenuRepository interface {
	ListMenu(ctx context.Context) ([]entity.Menu, error)
}
