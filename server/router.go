package server
import (
	"github.com/labstack/echo"
	"github.com/1bits.org/palabra/controladores"
)
func Router(e *echo.Group) {
	e.GET("/v1/hola",controladores.Hola )
	e.GET("/v1/id/:id",controladores.Presentarme)
	e.GET("/v1/verso/:id", controladores.Retornarid)
}
