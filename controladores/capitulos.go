package controladores

import (
	"net/http"

	"github.com/1-bits/palabra/modelo"
	"github.com/labstack/echo"
)

//get todos los libros
func GetListacapitulos(c echo.Context) (err error) {
	cap := new(modelo.Capitulo)
	if err = c.Bind(cap); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	cap.GetCountCapitulo()
	return c.JSON(http.StatusOK, cap)
}

//get todos los libros
func Getcapitulos(c echo.Context) (err error) {
	v := new(modelo.Versiculo)
	if err = c.Bind(v); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	resultado := new(modelo.ListaCapitulos)
	//
	return c.JSON(http.StatusOK, resultado.GetCapitulo(v))

}
