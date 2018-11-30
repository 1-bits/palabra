package controladores

import (
	"net/http"

	"github.com/1-bits/palabra/modelo"
	"github.com/labstack/echo"
)

//get todos los libros
func GetListacapitulos(c echo.Context) (err error) {
	v := new(modelo.Capitulo)
	if err = c.Bind(v); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	v.GetCapitulo()
	return c.JSON(http.StatusOK, v)
}
