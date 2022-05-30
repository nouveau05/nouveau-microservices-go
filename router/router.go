package router

import (
	"github.com/gorilla/mux"
	"github.com/nouveau05/nouveau-microservices-go/middleware"
)

// Router is exported and used in main.go

func Router() *mux.Router {

	router := mux.NewRouter()

	// Route information

	router.HandleFunc("/api/venture/{id}", middleware.GetVenture).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/venture", middleware.GetAllVentures).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newventure", middleware.CreateVenture).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/venture/{id}", middleware.UpdateVenture).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/venture/{id}", middleware.DeleteVenture).Methods("DELETE", "OPTIONS")

	return router
}
