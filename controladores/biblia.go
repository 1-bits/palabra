package controladores

import (
	"net/http"
	"strconv"

	"github.com/1bits.org/palabra/modelo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// prueba de conexion
func Versiculo(c echo.Context) error {
	versionstr := c.FormValue("version")
	version, err := strconv.Atoi(versionstr)
	if err != nil {
		// handle error
	}
	librostr := c.FormValue("libro")
	libro, err := strconv.Atoi(librostr)
	if err != nil {
		// handle error
	}
	capitulostr := c.FormValue("capitulo")
	capitulo, err := strconv.Atoi(capitulostr)
	if err != nil {
		// handle error
	}
	versostr := c.FormValue("verso")
	verso, err := strconv.Atoi(versostr)
	if err != nil {
		// handle error
	}
	result := modelo.GetVersiculo(version, libro, capitulo, verso)
	return c.JSON(http.StatusOK, result)
}
