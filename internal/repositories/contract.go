package repositories

import (
	"context"
)

// Storer store contract
type Storer interface {
	Store(ctx context.Context, param interface{}) (int64, error)
}

// Updater update contract
type Updater interface {
	Update(ctx context.Context, input interface{}, where interface{}) (int64, error)
}

// Deleter delete contract
type Deleter interface {
	Update(ctx context.Context, input interface{}, where interface{}) (int64, error)
}

// Counter count contract
type Counter interface {
	Count(ctx context.Context, p interface{}) (total int64, err error)
}
