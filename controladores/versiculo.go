package controladores

import (
	"net/http"

	"github.com/1bits.org/palabra/modelo"
	"github.com/labstack/echo"
)

// prueba de conexion
func RetornoVersiculo(c echo.Context) (err error) {
	v := new(modelo.Versiculo)
	if err = c.Bind(v); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	v.GetVersiculo()
	return c.JSON(http.StatusOK, v)
}
