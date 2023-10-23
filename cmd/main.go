package main

import (
	"context"
	"fmt"

	//defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"gitlab.com/tizim-back/api"
	_ "gitlab.com/tizim-back/api/docs"
	"gitlab.com/tizim-back/config"
	"gitlab.com/tizim-back/pkg/logger"
	"gitlab.com/tizim-back/storage"
)

func main() {
	cfg := config.Load(".")
	log := logger.New(cfg.LogLevel, "backend")

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.Postgres.PostgresHost),
		logger.String("port", cfg.Postgres.PostgresPort),
		logger.String("database", cfg.Postgres.PostgresDatabase),
	)

	connDB, err := pgxpool.New(context.Background(), cfg.Postgres.DatabaseURL)
	if err != nil {
		fmt.Println("failed connect database", err)
	}

	// // Load the model from a file.
	// m, err := model.NewModelFromFile(cfg.AuthConfigPath)
	// if err != nil {
	// 	panic(err)
	// }

	// // Load the policy from a CSV file.
	// casbinEnforcer, err := casbin.NewEnforcer(m, "./config/auth.csv")
	// if err != nil {
	// 	log.Error("Casbin conn")
	// 	fmt.Println(">>>>>>>>>>.", err)
	// }
	// err = casbinEnforcer.LoadPolicy()
	// if err != nil {
	// 	log.Error("casbin error load policy", logger.Error(err))
	// 	return
	// }

	//casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	//casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	strg := storage.NewStoragePg(connDB)

	apiServer := api.New(api.RoutetOptions{
		Cfg:     &cfg,
		Storage: strg,
		Log:     log,
	})

	if err = apiServer.Run(fmt.Sprintf(":%s", cfg.HttpPort)); err != nil {
		log.Fatal("failed to run server: %s", logger.Error(err))
	}
}
