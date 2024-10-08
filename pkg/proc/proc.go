package proc

import (
	"fmt"

	"github.com/deveshmishra34/groot/pkg/api/routers"
	clientsPkg "github.com/deveshmishra34/groot/pkg/clients"
	"github.com/deveshmishra34/groot/pkg/clients/cors"
	"github.com/deveshmishra34/groot/pkg/clients/dbc"
	"github.com/deveshmishra34/groot/pkg/clients/gzip"
	"github.com/deveshmishra34/groot/pkg/clients/kratos"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/clients/service"
	"github.com/deveshmishra34/groot/pkg/config"
	"github.com/deveshmishra34/groot/pkg/db/models"
	"github.com/deveshmishra34/groot/pkg/utils"
)

func InitServiceEnv(serviceName string, version string) {
	config.SetServiceName(serviceName)
	config.SetServiceVersion(version)
	config.InitEnv()
	config.ResolveDevMode()
	config.InitFeatures()
	config.ResolveFlags()
	config.PrintEnvInEnvMode()
}

var clients []clientsPkg.IClient

func InitClients() {
	InitServiceClient()
	InitCorsClient()
	InitGzipClient()
	InitDbClient()
	InitOryKratosClient()
	// ...
}

func ConfigureClients() {
	logger.Debug("Configuring clients ...")
	for _, c := range clients {
		feature := config.Feature(c.Name())
		if feature.IsEnabled() {
			logger.Debug("Configuring %s client ...", c.Name())
			c.Configure(feature.Config)
			continue
		}
		logger.Warn("Client: '%s' is disabled, This may cause runtime errors if this client is used.", c.Name())
	}
}

func addClient(client clientsPkg.IClient) {
	clients = append(clients, client)
}

func InitServiceClient() {
	client := service.GetClient()
	logger.Debug("Activating %s client ...", client.Name())
	addClient(client)
}

func InitCorsClient() {
	client := cors.GetClient()
	logger.Debug("Activating %s client ...", client.Name())
	addClient(client)
}

func InitGzipClient() {
	client := gzip.GetClient()
	logger.Debug("Activating %s client ...", client.Name())
	addClient(client)
}

func InitDbClient() {
	client := dbc.GetDBClient()
	logger.Debug("Activating %s client ...", client.Name())
	client.SetSilent(!config.DevModeFlag)
	addClient(client)
}

func InitOryKratosClient() {
	client := kratos.GetClient()
	logger.Debug("Activating %s client ...", client.Name())
	addClient(client)
}

func InitDbConnection() {
	logger.Debug("Initializing database connection ...")
	dbc.GetDBClient().InitDBConnection()
}

func InitModels() {
	logger.Debug("Activating models ...")
	models.Init(dbc.GetDBClient().DB)
}

func PrintHiddenRoutesTable() {
	routers.InitHiddenAPIRouter()
	routes := routers.HiddenAPIRouter().Echo.Routes()

	t := utils.PrepareRoutesTable(routes, "Hidden API Routes")
	utils.SetTableBorderStyle(t, config.NoBorderFlag)

	fmt.Println(t.Render())
}

func PrintProtectedRoutesTable() {
	routers.InitProtectedAPIRouter()
	routes := routers.ProtectedAPIRouter().Echo.Routes()

	t := utils.PrepareRoutesTable(routes, "Protected API Routes")
	utils.SetTableBorderStyle(t, config.NoBorderFlag)

	fmt.Println(t.Render())
}

func PrintPublicRoutesTable() {
	routers.InitPublicAPIRouter()
	routes := routers.PublicAPIRouter().Echo.Routes()

	t := utils.PrepareRoutesTable(routes, "Public API Routes")
	utils.SetTableBorderStyle(t, config.NoBorderFlag)

	fmt.Println(t.Render())
}
