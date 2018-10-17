package controladores

import (
	"net/http"

	"github.com/labstack/echo"
)

func Informacionversiculo(c echo.Context) error {

	return c.JSON(http.StatusOK, " assad ")
}
