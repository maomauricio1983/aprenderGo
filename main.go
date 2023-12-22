package main

import (
	"PrimerProyecto/rutas"
	"fmt"
	mux "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := mux.NewRouter()
	//rutas
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros) // ruta con parametros
	mux.HandleFunc("/parametros-query-string", rutas.ParametrosQueryString)

	// ejecucion del servidor
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}

	server := &http.Server{
		Addr:         "localhost:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("corriendo servidor en http://localhost:" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}

/*func main() {

	//mux := http.NewServeMux()

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(response, "hola mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}*/
