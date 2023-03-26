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

func (repo *OrderRepositoryImpl) Order(ctx context.Context, param interface{}) ([]entity.OrderResponse, error) {
	var (
		params        = param.(entity.OrderRequest)
		orderResponse []entity.OrderResponse
	)

	query := `INSERT INTO orders (privy_id, menu_id, quantity, price, status) VALUES($1, $2, $3, $4, $5) RETURNING id, status`
	price := entity.GetPrice(int32(params.MenuID))
	params.Status = entity.PROCESS

	err := repo.sql.Query(ctx, &orderResponse, query, params.PrivyID, params.MenuID, params.Quantity, price, params.Status)
	if err != nil {
		return orderResponse, err
	}
	return orderResponse, nil
}
