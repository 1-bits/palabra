package server

import (
	"net/http"

	"github.com/1-bits/palabra/controladores"
	"github.com/labstack/echo"
)

func Router(e *echo.Group) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/test/", controladores.Test)

	e.POST("/v1/libros/", controladores.GetLibro)
	e.GET("/v1/libros/", controladores.GetListaLibro)

	e.POST("/v1/versiones/", controladores.GetVersiones)
	e.GET("/v1/versiones/", controladores.GetListaVersiones)

	e.POST("/v1/capitulo/cantidad/", controladores.GetListacapitulos)
	e.POST("/v1/capitulo/", controladores.Getcapitulos)

	e.POST("/v1/versiculo/", controladores.RetornoVersiculo)

	e.POST("/v1/versiculo/", controladores.RetornoVersiculo)

}
