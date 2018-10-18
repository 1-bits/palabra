package controladores

import (
	"fmt"
	"net/http"

	"github.com/1bits.org/palabra/modelo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// prueba de conexion
func RetornoVersiculo(c echo.Context) (err error) {
	v := new(Versiculo)
	if err = c.Bind(v); err != nil {
		return c.JSON(http.StatusOK, "ERROR parametros")
	}
	result := modelo.GetVersiculo(v.version, v.libro, v.capitulo, v.verso)
	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}
