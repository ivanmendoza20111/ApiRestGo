package main

import (
	"ProyectoWeb/migration"
	"ProyectoWeb/routes"
	"flag"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la Migración a la Base de Datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la Migración...")
		migration.Migrate()
		log.Println("Finalizó la Migración.")
	}

	// Inicia las Rutas
	router := routes.InitRoutes()

	// Inicia los Middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	//Inicia el Servidor
	server := &http.Server{
		Addr: ":8080",
		Handler: n,
	}

	log.Println("Iniciando el Servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")

}
