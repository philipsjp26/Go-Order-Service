package menu

import (
	"context"

	"gitlab.privy.id/order_service/internal/entity"
	"gitlab.privy.id/order_service/pkg/databasex"
)

type MenuRepositoryImpl struct {
	sql databasex.Adapter
}

func NewMenuRepository(db databasex.Adapter) MenuRepository {
	return &MenuRepositoryImpl{sql: db}
}

func (repo *MenuRepositoryImpl) ListMenu(ctx context.Context) ([]entity.Menu, error) {
	var (
		menuResponse []entity.Menu
	)

	query := `SELECT * FROM menus`

	err := repo.sql.Query(ctx, &menuResponse, query)
	if err != nil {
		return menuResponse, err
	}
	return menuResponse, nil
}
