package modelo

import (
	"github.com/1-bits/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

//Versiculo
type (
	ContextVersiculo interface {
		GetVersiculo()
	}
	Versiculo struct {
		Version  int    `json: "version" form:"version" query:"version"`
		Libro    int    `json: "libro" form:"libro" query:"libro"`
		Capitulo int    `json: "capitulo" form:"capitulo" query:"capitulo"`
		Verso    int    `json: "verso" form:"verso" query:"verso"`
		Texto    string `json: "texto" form:"texto" query:"texto"`
	}
)

func (v *Versiculo) GetVersiculo() {
	con := db.CreateCon()
	err := con.QueryRow("SELECT Texto FROM Versos where Version=? and Libro=? and Capitulo=? and Verso=? ", v.Version, v.Libro, v.Capitulo, v.Verso).Scan(&v.Texto)
	if err != nil {
		v.Texto = "Lo sentimos este texto no existe"
	}
}
