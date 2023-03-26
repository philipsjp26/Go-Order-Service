// Package ucase
package ucase

import (
	"testing"

	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Serve(t *testing.T) {
	svc := NewHealthCheck()

	t.Run("test health check", func(t *testing.T) {
		result := svc.Serve(&appctx.Data{})

		assert.Equal(t, appctx.Response{
			Code:    consts.CodeSuccess,
			Message: "ok",
		}, result)
	})
}
