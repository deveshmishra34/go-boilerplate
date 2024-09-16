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
	"github.com/codoworks/go-boilerplate/pkg/db/models"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Delete(c echo.Context) error {
	// id, err := handlers.GetUUIDParam(c.Param("id"))
	// if err != nil {
	// 	c.Logger().Error(constants.ERROR_ID_NOT_FOUND)
	// 	return constants.ERROR_ID_NOT_FOUND
	// }
	// kratosCli := kratos.GetClient()
	// if err := kratosCli.DeleteIdentity(id.String()); err != nil {
	// 	return err
	// }
	// return c.JSON(http.StatusAccepted, handlers.Accepted())

	id := c.Param("id")
	if id == "" {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, nil)
	}

	if err := models.UserModel().Delete(id); err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusAccepted, handlers.Deleted())
}
