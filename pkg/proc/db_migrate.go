package proc

import (
	"github.com/deveshmishra34/groot/pkg/clients/dbc"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/db/migrations"
)

func DBMigrate() {
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitDBConnection()

	migrations.Init(dbClient.DB)

	if err := migrations.Migrate(); err != nil {
		logger.Error("Failed to apply migrations: %s", err)
	}

	logger.Info("Migrations applied successfully")

}
