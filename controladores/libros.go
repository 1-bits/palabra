package controladores

import (
	"net/http"

	"github.com/1-bits/palabra/modelo"
	"github.com/labstack/echo"
)

// prueba de conexion
func GetLibro(c echo.Context) (err error) {
	l := new(modelo.Libros)
	resultado := new(modelo.ListaLibros)
	if err = c.Bind(l); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	if l.Id != 0 {
		a := l.GetNombreLibro()
		resultado.Lista = a.Lista
	} else {
		a := l.BuscarIdNombre()
		resultado.Lista = a.Lista
	}
	return c.JSON(http.StatusOK, resultado)
}

//get todos los libros
func GetListaLibro(c echo.Context) (err error) {
	resultado := new(modelo.ListaLibros)
	//	resultado.GetTodosLosLibros()
	return c.JSON(http.StatusOK, resultado.GetTodosLosLibros())
}
