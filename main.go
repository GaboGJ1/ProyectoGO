package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Crear un nuevo enrutador
	r := mux.NewRouter()

	// Manejar la ruta principal "/"
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Manejar la ruta "/env"
	r.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, ENV! " + os.Getenv("TEST_ENV")))
	})

	// Servir archivos estáticos desde el directorio "./static"
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Obtener el puerto desde la variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Escuchando en el puerto %s\n", port)
	err := http.ListenAndServe("0.0.0.0:"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
