package healthz

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/config"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	payload := map[string]string{
		"message": "ok",
		"version": config.Env.Version,
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))
}
