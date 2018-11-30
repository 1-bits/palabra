package controladores

import (
	"net/http"

	"github.com/1-bits/palabra/modelo"
	"github.com/labstack/echo"
)

//get todos los libros
func GetListacapitulos(c echo.Context) (err error) {
	ca := new(modelo.Capitulo)
	if err = c.Bind(ca); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	resultado := new(modelo.Listacapitulos)
	return c.JSON(http.StatusOK, resultado.GetcapituloPorLibro(ca.Capitulo))
}
