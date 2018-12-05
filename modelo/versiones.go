package modelo

import (
	"github.com/1-bits/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

type Versiones struct {
	Id          int    `json: "id" form:"id" query:"id"`
	Version     string `json: "version" form:"version" query:"version"`
	Descripcion string `json: "descripcion" form:"descripcion" query:"descripcion"`
}
type ListaVersiones struct {
	Lista []Versiones `json: "Versiones" form:"Versiones" query:"Versiones"`
}

func (v *Versiones) GetVersiculo() {
	con := db.CreateCon()
	err := con.QueryRow("SELECT Texto FROM Versos where Version=? and Libro=? and Capitulo=? and Verso=? ").Scan()
	defer con.Close()
	if err != nil {

		//&v.Texto = "Lo sentimos este texto no existe "
	}
}

func (l *ListaVersiones) GetTodoslasversiones() ListaVersiones {
	ver := new(ListaVersiones)
	con := db.CreateCon()
	rows, err := con.Query("SELECT id,version,descripcion FROM version")
	defer rows.Close()
	defer con.Close()
	for rows.Next() {
		Version := Versiones{}
		err = rows.Scan(&Version.Id, &Version.Version, &Version.Descripcion)
		if err != nil {
			panic(err)
		}
		ver.Lista = append(ver.Lista, Version)
	}
	err = rows.Err()
	if err != nil {
		panic(err)

	}
	return *ver
}
