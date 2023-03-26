// Package bootstrap
package bootstrap

import (
	"time"
	"fmt"

	"gitlab.privy.id/order_service/pkg/databasex"
	"gitlab.privy.id/order_service/pkg/logger"
	"gitlab.privy.id/order_service/internal/appctx"
)

// RegistryDatabase initialize database session
func RegistryDatabase(cfg *appctx.Database) *databasex.DB {

	db, err := databasex.CreateSession(&databasex.Config{
		Driver:       cfg.Driver,
		Host:         cfg.Host,
		Name:         cfg.Name,
		Password:     cfg.Pass,
		Port:         cfg.Port,
		User:         cfg.User,
		Timeout:      cfg.TimeoutSecond,
		MaxOpenConns: cfg.MaxOpen,
		MaxIdleConns: cfg.MaxIdle,
		MaxLifetime:  time.Duration(cfg.MaxLifeTimeMS) * time.Millisecond,
		Charset:      cfg.Charset,
		TimeZone:     cfg.Timezone,
	})
	if err != nil {
		logger.Fatal(
			err,
			logger.EventName("db"),
			logger.Any("host", cfg.Host),
			logger.Any("port", cfg.Port),
			logger.Any("driver", cfg.Driver),
			logger.Any("timezone", cfg.Timezone),
		)
	}
	//db := databasex.New()
	return databasex.New(db, false, cfg.Name)
}

// RegistryMultiDatabase initialize database session
func RegistryMultiDatabase(cfgWrite *appctx.Database, cfgRead *appctx.Database) databasex.Adapter {
	lf := logger.NewFields(
		logger.EventName("db"),
		logger.Any("host_read", cfgRead.Host),
		logger.Any("port_read", cfgRead.Port),
		logger.Any("host_write", cfgWrite.Host),
		logger.Any("port_write", cfgWrite.Port),
		logger.Any("driver_write", cfgWrite.Driver),
		logger.Any("timezone_write", cfgWrite.Timezone),
		logger.Any("driver_read", cfgRead.Driver),
		logger.Any("timezone_read", cfgRead.Timezone),
	)
	dbWrite, err := databasex.CreateSession(&databasex.Config{
		Driver:       cfgWrite.Driver,
		Host:         cfgWrite.Host,
		Name:         cfgWrite.Name,
		Password:     cfgWrite.Pass,
		Port:         cfgWrite.Port,
		User:         cfgWrite.User,
		Timeout:      cfgWrite.TimeoutSecond,
		MaxOpenConns: cfgWrite.MaxOpen,
		MaxIdleConns: cfgWrite.MaxIdle,
		MaxLifetime:  time.Duration(cfgWrite.MaxLifeTimeMS) * time.Millisecond,
		Charset:      cfgWrite.Charset,
		TimeZone:     cfgWrite.Timezone,
	})

	if err != nil {
		logger.Fatal(fmt.Sprintf("db write %v", err), lf...)
	}

	dbRead, err := databasex.CreateSession(&databasex.Config{
		Driver:       cfgWrite.Driver,
		Host:         cfgWrite.Host,
		Name:         cfgWrite.Name,
		Password:     cfgWrite.Pass,
		Port:         cfgWrite.Port,
		User:         cfgWrite.User,
		Timeout:      cfgWrite.TimeoutSecond,
		MaxOpenConns: cfgWrite.MaxOpen,
		MaxIdleConns: cfgWrite.MaxIdle,
		MaxLifetime:  time.Duration(cfgWrite.MaxLifeTimeMS) * time.Millisecond,
		Charset:      cfgWrite.Charset,
		TimeZone:     cfgWrite.Timezone,
	})

	if err != nil {
		logger.Fatal(fmt.Sprintf("db read %v", err), lf...)
	}

	return databasex.NewMulti(databasex.New(dbWrite, false, cfgRead.Name), databasex.New(dbRead, true, cfgRead.Name))
}


