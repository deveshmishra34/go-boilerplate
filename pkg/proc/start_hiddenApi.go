package proc

import (
	"github.com/deveshmishra34/groot/pkg/api/routers"
	"github.com/deveshmishra34/groot/pkg/clients/service"
)

func StartHiddenApi() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	routers.InitHiddenAPIRouter()
	routers.HiddenAPIRouter().Start(config.Host, config.HiddenApiPort)
}
