package genx

import (
	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/pkg/structgen"
)

func Gen() {
	cfg := appctx.NewConfig()
	structgen.CreateAll(structgen.Configuration{
		DbHost:     cfg.WriteDB.Host,
		DbName:     cfg.WriteDB.Name,
		DbUser:     cfg.WriteDB.User,
		DbPassword: cfg.WriteDB.Pass,
		TagLabel:   "db,json",
		Driver:     cfg.WriteDB.Driver,
		Timezone:   cfg.WriteDB.Timezone,
		DbPort:     cfg.WriteDB.Port,
	})
}

func GenLogic() {
	structgen.CreateLogic()
}

func GenEntity() {
	cfg := appctx.NewConfig()
	structgen.CreateEntity(structgen.Configuration{
		DbHost:     cfg.WriteDB.Host,
		DbName:     cfg.WriteDB.Name,
		DbUser:     cfg.WriteDB.User,
		DbPassword: cfg.WriteDB.Pass,
		TagLabel:   "db,json",
		Driver:     cfg.WriteDB.Driver,
		Timezone:   cfg.WriteDB.Timezone,
		DbPort:     cfg.WriteDB.Port,
	})
}

func GenPresentation() {
	cfg := appctx.NewConfig()
	structgen.CreatePresentation(structgen.Configuration{
		DbHost:     cfg.WriteDB.Host,
		DbName:     cfg.WriteDB.Name,
		DbUser:     cfg.WriteDB.User,
		DbPassword: cfg.WriteDB.Pass,
		TagLabel:   "db,json",
		Driver:     cfg.WriteDB.Driver,
		Timezone:   cfg.WriteDB.Timezone,
		DbPort:     cfg.WriteDB.Port,
	})
}


