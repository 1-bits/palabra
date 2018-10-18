package modelo

import (
	"github.com/1bits.org/palabra/db"
	_ "github.com/go-sql-driver/mysql"
)

func GetVersiculo(version int, libro int, capitulo int, versiculo int) Versiculo {
	resultado := Versiculo{}
	con := db.CreateCon()
	err := con.QueryRow("SELECT version, libro,	capitulo, verso, texto FROM versos where version=? and libro=? and capitulo=? and verso=? ", version, libro, capitulo, versiculo).Scan(&resultado.version, &resultado.libro, &resultado.capitulo, &resultado.verso, &resultado.texto)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	//var resultado Versiculos
	//resultado.Versiculo = append(resultado.Versiculo, resultado)
	return resultado
}
