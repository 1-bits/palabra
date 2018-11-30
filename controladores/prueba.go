package controladores

import (
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, "hola")
}
