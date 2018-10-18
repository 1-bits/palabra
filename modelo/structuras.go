package modelo

//Versiculo
type Versiculo struct {
	version  int    `json: "verso_version" form:"version"`
	libro    int    `json: "verso_libro" form:"libro"`
	capitulo int    `json: "verso_capitulo" form:"capitulo"`
	verso    int    `json: "verso_versiculoid" form:"verso"`
	texto    string `json: "verso_texto" form:"texto"`
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
