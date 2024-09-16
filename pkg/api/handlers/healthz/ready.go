package healthz

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/api/helpers"
	"github.com/deveshmishra34/groot/pkg/clients/dbc"

	"github.com/labstack/echo/v4"
)

func Ready(c echo.Context) error {
	dbClient := dbc.GetDBClient()
	if err := dbClient.Ping(); err != nil {
		return helpers.Error(c, err, nil)
	}

	payload := map[string]string{
		"message": "ready",
	}
	return c.JSON(http.StatusOK, handlers.Success(payload))
}
