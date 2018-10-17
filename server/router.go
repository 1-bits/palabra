package server

import (
	"github.com/1bits.org/palabra/controladores"
	"github.com/labstack/echo"
)

func Router(e *echo.Group) {
	e.POST("/v1/versiculo/", controladores.Versiculo)
	e.GET("/v1/versiculo/", controladores.Informacionversiculo)
	//	e.POST("/v1/capitulo/", controladores.Retornarid)
}
