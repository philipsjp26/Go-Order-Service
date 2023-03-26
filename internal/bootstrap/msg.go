// Package bootstrap
package bootstrap

import (
	"gitlab.privy.id/order_service/internal/consts"
	"gitlab.privy.id/order_service/pkg/logger"
	"gitlab.privy.id/order_service/pkg/msgx"
)

func RegistryMessage()  {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
