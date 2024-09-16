package errors

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context) error {
	return c.JSON(
		http.StatusNotFound,
		handlers.BuildResponse(
			constants.STATUS_CODE_ROUTE_NOT_FOUND,
			constants.MSG_ROUTE_NOT_FOUND,
			[]string{},
			nil))
}
