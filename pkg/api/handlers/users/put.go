package users

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/api/helpers"
	"github.com/deveshmishra34/groot/pkg/db/models"
	"github.com/deveshmishra34/groot/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, nil)
	}

	f := &models.UserForm{}

	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if err := helpers.Validate(f); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ValidationErrors(err))
	}

	m := f.MapToModel()

	m.ID = id

	if err := m.Update(); err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))

}
