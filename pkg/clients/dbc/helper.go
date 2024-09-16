package dbc

import (
	"github.com/deveshmishra34/groot/pkg/clients/dbc/adapters"
	"github.com/deveshmishra34/groot/pkg/utils/constants"

	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

var dbClient *DBClient

func init() {
	dbClient = &DBClient{
		name:    constants.FEATURE_DATABASE,
		adapter: adapters.Adapters,
		silent:  true,
		gormConfig: &gorm.Config{
			Logger: gLogger.Default.LogMode(gLogger.Silent),
		},
	}
}

func GetDBClient() *DBClient {
	return dbClient
}
