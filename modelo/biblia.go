package modelo

import (
	"fmt"

	"github.com/1bits.org/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

func GetVersiculo(version int, libro int, capitulo int, versiculo int) Versiculo {
	fmt.Println(version)
	fmt.Println(libro)
	fmt.Println(capitulo)
	fmt.Println(versiculo)
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM palabradb.versos where version=1 and libro=1 and capitulo=1 and verso=1 "
	rows, err := con.Query(sqlStatement) //, version, libro, capitulo, versiculo)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	selectVersiculos := Versiculo{}
	for rows.Next() {
		err := rows.Scan(&selectVersiculos.version, &selectVersiculos.libro, &selectVersiculos.capitulo,
			&selectVersiculos.verso,
			&selectVersiculos.texto)
		if err != nil {
		}
	}
	return selectVersiculos

}
