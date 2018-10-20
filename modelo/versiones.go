package modelo

import (
	"github.com/1-bits/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

type Versiones struct {
	Id          int    `json: "id" form:"id" query:"id"`
	Version     string `json: "version" form:"version" query:"version"`
	descripcion string `json: "descripcion" form:"descripcion" query:"descripcion"`
}

func (v *Versiones) GetVersiculo() {
	con := db.CreateCon()
	err := con.QueryRow("SELECT Texto FROM Versos where Version=? and Libro=? and Capitulo=? and Verso=? ").Scan()
	if err != nil {
		//&v.Texto = "Lo sentimos este texto no existe "
	}
}
