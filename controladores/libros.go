package controladores

import (
	"fmt"
	"net/http"

	"github.com/1bits.org/palabra/modelo"
	"github.com/labstack/echo"
)

// prueba de conexion
func GetLibro(c echo.Context) (err error) {
	l := new(modelo.Libros)
	resultado := new(modelo.Libross)
	if err = c.Bind(l); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	if l.Id != 0 {
		a := l.GetNombreLibro()
		fmt.Println(a)
		resultado.Lista = a.Lista
	} else {
		//a := new(modelo.Libross)
		a := l.BuscarIdNombre()
		fmt.Println(a)
		resultado.Lista = a.Lista
	}

	return c.JSON(http.StatusOK, resultado)
}
