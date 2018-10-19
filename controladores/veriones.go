package controladores

import (
	"net/http"

	"github.com/labstack/echo"
)

// prueba de conexion
func GetVersiones(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, "hola")
}
