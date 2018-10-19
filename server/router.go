package server

import (
	"github.com/1bits.org/palabra/controladores"
	"github.com/labstack/echo"
)

func Router(e *echo.Group) {
	e.POST("/v1/versiculo/", controladores.RetornoVersiculo)
	e.POST("/v1/libros/", controladores.GetLibro)
	e.POST("/v1/versiones/", controladores.GetVersiones)
	e.POST("/v1/capitulo/", controladores.GetCapitulos)
}
