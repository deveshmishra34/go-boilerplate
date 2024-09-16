package middlewares

import (
	"github.com/deveshmishra34/groot/pkg/clients/gzip"
	"github.com/deveshmishra34/groot/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GzipMiddleware() echo.MiddlewareFunc {
	GzipCli := gzip.GetClient()
	config := GzipCli.GetConfig()
	level := utils.IntFromStr(config.Level)
	return middleware.GzipWithConfig(middleware.GzipConfig{Level: level})
}
