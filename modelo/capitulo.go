package modelo

import (
	"github.com/1-bits/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

type (
	ContextCapitulo interface {
		GetCapitulo()
	}
	Capitulo struct {
		Libro    int `json: "Libro" form:"Libro" query:"Libro"`
		Capitulo int `json: "Capitulo" form:"Capitulo" query:"Capitulo"`
	}
)

func (ca *Capitulo) GetCapitulo() {
	con := db.CreateCon()
	err := con.QueryRow("SELECT count(capitulo) as Capitulo FROM versos WHERE version=1 and libro=? and texto<>'' GROUP by capitulo ORDER BY versos.capitulo ", ca.Libro).Scan(&ca.Capitulo)
	if err != nil {
		ca.Capitulo = 0
	}
}
