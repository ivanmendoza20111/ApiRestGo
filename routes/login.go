package routes

import (
	"ProyectoWeb/controllers"
	"github.com/gorilla/mux"
)

// SetLogin router para Login
func SetLoginRouter(router *mux.Router){
	router.HandleFunc("/api/login",controllers.Login).Methods("POST")
}
