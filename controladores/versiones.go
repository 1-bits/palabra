package controladores

import (
	"net/http"

	"github.com/1-bits/palabra/modelo"
	"github.com/labstack/echo"
)

// prueba de conexion
func GetVersiones(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, "hola")
}

//get todos los libros
func GetListaVersiones(c echo.Context) (err error) {
	resultado := new(modelo.ListaVersiones)
	//	resultado.GetTodosLosLibros()
	return c.JSON(http.StatusOK, resultado.GetTodoslasversiones())
}
