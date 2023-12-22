package rutas

import (
	"fmt"
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

// recibiendo argumentos enviados en la url
func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID= "+vars["id"]+"  | SLUG= "+vars["slug"])
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, request.URL)                     // devuelve toda la url:  /parametros-query-string?id=78&slug=mauricio
	fmt.Fprintln(response, request.URL.RawQuery)            // devuelve los parametros : id=78&slug=mauricio
	fmt.Fprintln(response, request.URL.Query())             // devuelve los parametros como un mapa : map[id:[78] slug:[mauricio]]
	fmt.Fprintln(response, request.URL.Query().Get("id"))   // devuelve el parametro solicitado: id
	fmt.Fprintln(response, request.URL.Query().Get("slug")) // devuelve el parametro solicitado: slug

	id := request.URL.Query().Get("id") // lo guardamos en una variable id
	slug := request.URL.Query().Get("slug")
	fmt.Fprintln(response, "id= %s | slug= %s", id, slug) // lo guardamos en una variable slug
}
