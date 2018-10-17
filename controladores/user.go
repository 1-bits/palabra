package controladores

import (
	"net/http"

	"github.com/labstack/echo"
)

func Presentarme(c echo.Context) error {

	return c.JSON(http.StatusOK, "FUNCIONA")
}
