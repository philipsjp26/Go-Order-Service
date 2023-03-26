package order

import (
	"context"

	"gitlab.privy.id/order_service/internal/entity"
)

type OrderRepository interface {
	Order(ctx context.Context, param interface{}) (entity.OrderResponse, error)
}
