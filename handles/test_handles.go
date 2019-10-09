package handles

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/tranhuy-dev/K8s/models"
)
func TestHandles(c echo.Context) error {
	return c.JSON(http.StatusOK , models.TestModels{
		Code: 200,
		Message: "Success",
	})
}