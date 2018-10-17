package modelo

//Versiculo
type Versiculo struct {
	version  string `json: "verso_version"`
	libro    string `json: "verso_libro"`
	capitulo string `json: "verso_capitulo"`
	verso    string `json:"verso_versiculoid"`
	texto    string `json: "verso_texto"`
}

//Capitulos
type Capitulo struct {
	version  string `json: "verso_version"`
	libro    string `json: "verso_libro"`
	capitulo string `json: "verso_capitulo"`
	texto    string `json: "verso_texto"`
}

//usuario
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
