package controladores

import (
	"github.com/labstack/echo"
	"fmt"
	"database/sql"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

// prueba de conexion
func Retornarid(c echo.Context) error {
	requested_id := c.Param("id")
	fmt.Println(requested_id);
	db, err := sql.Open("mysql", "analistas:analistas@tcp(acapdes01d.asociacioncibao.com.do:3306)/go_prueba")
	if err != nil {
		fmt.Println(err.Error())
		response := Excuse{Id: "", Error: "conection error ", Quote: ""}
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer db.Close()
	var quote string;
	var id string;
	err = db.QueryRow("SELECT id, quote FROM excuses WHERE id = ?", requested_id).Scan(&id, &quote)
	if err != nil {
		fmt.Println(err)
	}

	libro := Libros{Name:id+quote, Email:quote}
	response := Excuse{Id: id, Error: "null", Quote: quote,Libro:libro}
	return c.JSON(http.StatusOK, response)
}
