package main

import (
	"fmt"

	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"

	"github.com/casbin/casbin/v2"
	"gitlab.com/backend/api"
	"gitlab.com/backend/config"
	"gitlab.com/backend/pkg/db"
	"gitlab.com/backend/pkg/logger"
	"gitlab.com/backend/storage"
)

func main() {

	fmt.Println("Server is running!!!")

	cfg := config.Load()
	log := logger.New("debug", "clinic")

	var (
		casbinEnforcer *casbin.Enforcer
	)

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CsvFilePath)
	if err != nil {
		log.Error("casbin enforcer error: ", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	connDb, err := db.ConnectToDb(cfg)

	if err != nil {
		log.Fatal("Couldn't connect to database: ", logger.Error(err))
	}

	strg := storage.NewStoragePg(connDb)

	apiServer := api.New(&api.Options{
		Cfg:            cfg,
		Storage:        strg,
		Log:            log,
		CasbinEnforcer: casbinEnforcer,
	})

	err = apiServer.Run(":" + cfg.HttpPort)
	if err != nil {
		log.Fatal("failed to run server: %v", logger.Error(err))
	}
	log.Fatal("Server Stopped")
}
