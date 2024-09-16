package cats

import (
	"net/http"

	"github.com/deveshmishra34/groot/pkg/api/handlers"
	"github.com/deveshmishra34/groot/pkg/api/helpers"
	"github.com/deveshmishra34/groot/pkg/db/models"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	ms, err := models.CatModel().FindAll()

	if err != nil {
		return helpers.Error(c, err, nil)
	}

	var payload []*models.CatForm

	for _, m := range ms {
		f := m.MapToForm()
		payload = append(payload, f)
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))

}
