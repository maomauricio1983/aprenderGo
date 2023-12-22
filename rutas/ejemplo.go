package rutas

import (
	"github.com/gorilla/mux"
	template2 "html/template"
	"net/http"
)

// carga una plantilla  home.html
func Home(response http.ResponseWriter, request *http.Request) {
	template, err := template2.ParseFiles("templates/ejemplo/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

/*func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola mundo desde Golang")
}*/

// carga una plantilla  nosotros.html
func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err := template2.ParseFiles("templates/ejemplo/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

/*func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "pagina sobre nosotros")
	fmt.Println("algo desde la terminal")
}*/

// recibiendo argumentos enviados en la url ->    http://localhost:8080/parametros/78/mauricio
func Parametros(response http.ResponseWriter, request *http.Request) {
	template, err := template2.ParseFiles("templates/ejemplo/parametros.html")
	vars := mux.Vars(request)
	texto := "prueba de parametros"
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": texto,
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}

// recibe parametros desde la url     http://localhost:8080/parametros-query-string?id=78&slug=mauricio
func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	template, err := template2.ParseFiles("templates/ejemplo/parametros_querystring.html")
	data := map[string]string{
		"id":   request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}
