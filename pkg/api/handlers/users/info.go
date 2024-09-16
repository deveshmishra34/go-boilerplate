package users

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/api/helpers"
	"github.com/deveshmishra34/groot/pkg/clients/logger"
	"github.com/deveshmishra34/groot/pkg/db/models"

	"github.com/labstack/echo/v4"
)

func Info(c echo.Context) error {
	uID := c.Request().Header.Get("User-Claim-Issuer")
	logger.Debug("User-Claim-Issuer: %+v", uID)

	user, err := models.UserModel().Find(uID)
	if err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(user.MapToInfoForm()))
}
