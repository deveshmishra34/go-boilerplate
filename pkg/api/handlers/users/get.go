package users

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/api/helpers"
	"github.com/deveshmishra34/groot/pkg/db/models"
	"github.com/deveshmishra34/groot/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, nil)
	}

	m, err := models.UserModel().Find(id)
	if err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))

	// id, err := handlers.GetUUIDParam(c.Param("id"))
	// if err != nil {
	// 	c.Echo().Logger.Error(constants.ERROR_ID_NOT_FOUND)
	// 	return constants.ERROR_ID_NOT_FOUND
	// }
	// kratosCli := kratos.GetClient()
	// identity, err := kratosCli.GetIdentity(id.String())
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(http.StatusOK, handlers.Success(identity))
}
