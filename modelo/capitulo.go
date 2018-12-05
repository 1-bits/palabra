package modelo

import (
	"github.com/1-bits/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

type (
	ContextCapitulo interface {
		GetCountCapitulo()
	}
	Capitulo struct {
		Libro    int `json: "Libro" form:"Libro" query:"Libro"`
		Capitulo int `json: "Capitulo" form:"Capitulo" query:"Capitulo"`
	}
)

type ListaCapitulos struct {
	Lista []Versiculo `json: "Versiculo" form:"Versiculo" query:"Versiculo"`
}

func (ca *Capitulo) GetCountCapitulo() {
	con := db.CreateCon()
	err := con.QueryRow("SELECT count(capitulo) as Capitulo FROM versos WHERE version=1 and libro=? and texto<>'' GROUP by capitulo ORDER BY versos.capitulo ", ca.Libro).Scan(&ca.Capitulo)
	defer con.Close()
	if err != nil {
		defer con.Close()
		ca.Capitulo = 0
	}
}

func (l *ListaCapitulos) GetCapitulo(v Versiculo) ListaCapitulos {
	lis := new(ListaCapitulos)
	con := db.CreateCon()
	rows, err := con.Query("SELECT version,libro,capitulo,verso,texto FROM versos WHERE version=? and libro=? and capitulo=?", v.Version, v.Libro, v.Capitulo)
	defer rows.Close()
	defer con.Close()
	for rows.Next() {
		Versiculo := Versiculo{}
		err = rows.Scan(&Versiculo.Version, &Versiculo.Libro, &Versiculo.Capitulo, &Versiculo.Verso, &Versiculo.Texto)
		if err != nil {
			panic(err)
		}
		lis.Lista = append(lis.Lista, Versiculo)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return *lis
}
