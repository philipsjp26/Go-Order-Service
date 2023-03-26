// Package contract
package contract

import (
	"context"

	"gitlab.privy.id/order_service/internal/appctx"
)


// UseCase is a use case contract
type UseCase interface {
	Serve(data *appctx.Data) appctx.Response
}


// MessageProcessor is use case queue message processor contract
type MessageProcessor interface {
	Serve(ctx context.Context, data *appctx.ConsumerData) error
}
