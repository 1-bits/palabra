package modelo

import (
	"fmt"
	"strings"

	"github.com/1bits.org/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

type Libros struct {
	Id     int    `json: "version" form:"version" query:"version"`
	Nombre string `json: "nombre" form:"nombre" query:"nombre"`
}

type Libross struct {
	Lista []Libros `json: "Libros" form:"Libros" query:"Libros"`
}

//)
func (l *Libros) GetNombreLibro() Libross {
	lis := new(Libross)
	con := db.CreateCon()
	err := con.QueryRow("SELECT nombre FROM libros where id=? ", l.Id).Scan(&l.Nombre)
	if err != nil {
		//&v.Texto = "Lo sentimos este texto no existe "
	}
	lis.Lista = append(lis.Lista, *l)
	return *lis
}

//
func (l *Libros) BuscarIdNombre() Libross {
	lis := new(Libross)
	con := db.CreateCon()
	rows, err := con.Query("SELECT id, nombre FROM libros where UPPER(nombre) like ? ", strings.ToUpper(fmt.Sprint("%", l.Nombre, "%")))
	defer rows.Close()
	for rows.Next() {
		Libros := Libros{}
		err = rows.Scan(&Libros.Id, &Libros.Nombre)
		if err != nil {
			panic(err)
		}
		lis.Lista = append(lis.Lista, Libros)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return *lis
}
