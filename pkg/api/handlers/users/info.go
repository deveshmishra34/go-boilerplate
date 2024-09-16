/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package users

import (
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/api/helpers"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/db/models"

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
