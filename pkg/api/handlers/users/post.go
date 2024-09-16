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
	"github.com/codoworks/go-boilerplate/pkg/utils"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	// newUser := &models.UserObj{}
	// if err := c.Bind(newUser); err != nil {
	// 	c.Echo().Logger.Error(constants.ERROR_BINDING_BODY)
	// 	c.Echo().Logger.Error(err)
	// 	return constants.ERROR_BINDING_BODY
	// }
	// if err := Validate.Struct(newUser); err != nil {
	// 	r := handlers.BuildValidationErrorsResponse(err.(validator.ValidationErrors))
	// 	return c.JSON(http.StatusBadRequest, r)
	// }
	// identity, err := kratos.Cli.CreateIdentity(newUser.ConvertToMap(), newUser.Password)
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(http.StatusOK, handlers.Success(identity))

	f := &models.UserForm{}
	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}
	if err := helpers.Validate(f); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ValidationErrors(err))
	}

	// Todo: Changes this hashing algorithm to a more secure and optimise one
	// Hashing the password with a random salt
	hashedPassword, err := utils.HashPassword(f.Password)

	if err != nil {
		return helpers.Error(c, constants.ERROR_INTERNAL_SERVER, err)
	}
	f.Password = string(hashedPassword)

	m := f.MapToModel()

	if err := m.Save(); err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))
}
