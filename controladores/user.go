package controladores

import (
		"github.com/labstack/echo"
		"net/http"
)


func  Hola(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hi!")
}

func Presentarme(c echo.Context) error {
	usuario :=User{
		Name:"cesar",
		Email:"cesardarinel@prueba.com",
	}
	return c.JSON(http.StatusOK,usuario)
}