package order

import (
	"context"

	"gitlab.privy.id/order_service/internal/entity"
	"gitlab.privy.id/order_service/pkg/databasex"
)

type OrderRepositoryImpl struct {
	sql databasex.Adapter
}

func NewOrderRepository(db databasex.Adapter) OrderRepository {
	return &OrderRepositoryImpl{sql: db}
}

func (repo *OrderRepositoryImpl) Order(ctx context.Context, param interface{}) (entity.OrderResponse, error) {
	panic(nil)
}
