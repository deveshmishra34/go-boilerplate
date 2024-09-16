package proc

import (
	"github.com/deveshmishra34/groot/pkg/api/routers"
	"github.com/deveshmishra34/groot/pkg/clients/service"
)

func StartProtectedApi() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	routers.InitProtectedAPIRouter()
	routers.ProtectedAPIRouter().Start(config.Host, config.ProtectedApiPort)
}
