package middlewares

import (
	"time"

	"github.com/deveshmishra34/groot/pkg/clients/service"
	"github.com/deveshmishra34/groot/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TimeoutMiddleware() echo.MiddlewareFunc {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	timeoutDuration := utils.IntFromStr(config.RequestTimeoutDuration)

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Duration(timeoutDuration) * time.Second,
	})
}
