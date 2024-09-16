package proc

import (
	"github.com/deveshmishra34/groot/pkg/api/routers"
	"github.com/deveshmishra34/groot/pkg/clients/service"
)

func StartPublicApi() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	routers.InitPublicAPIRouter()
	routers.PublicAPIRouter().Start(config.Host, config.PublicApiPort)
}
