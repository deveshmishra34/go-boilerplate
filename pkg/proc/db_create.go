package proc

import (
	"fmt"

	"github.com/deveshmishra34/groot/pkg/clients/dbc"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
)

func DBCreate() {
	// init feature [database]
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitServerConnection()

	if err := dbClient.CreateDatabase(); err != nil {
		panic(fmt.Errorf("failed to create database: %w", err))
	}

	// logger.Info("Database '" + config.Env.Config.DBName + "' created successfully.")

}
