package proc

import (
	"github.com/deveshmishra34/groot/pkg/clients/dbc"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/db/seeds"
)

func DBSeed() {
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitDBConnection()

	seeds.Init(dbClient.DB)

	if err := seeds.Apply(); err != nil {
		logger.Error("Failed to apply seeds: %s", err)
	}

	logger.Info("Seeds applied successfully")

}
