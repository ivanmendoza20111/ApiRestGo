package main

import (
	"ProyectoWeb/migration"
	"flag"
	"log"
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
}
