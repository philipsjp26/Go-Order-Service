// Package bootstrap
package bootstrap

import (
	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/pkg/logger"
	"gitlab.privy.id/order_service/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.Name,
	})
}

