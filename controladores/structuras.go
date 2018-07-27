package controladores
// estructura de control
type Excuse struct {
	Error string `json:"error"`
	Id string `json:"id"`
	Quote string `json:"quote"`
	Libro Libros
}
//libros
type Libros struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
//usuario
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
