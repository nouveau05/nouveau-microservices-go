package router

import (
	"go-nouveau-postgres-api/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/venture/{id}", middleware.GetVenture).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/venture", middleware.GetAllVentures).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newventure", middleware.CreateVenture).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/venture/{id}", middleware.UpdateVenture).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/venture/{id}", middleware.DeleteVenture).Methods("DELETE", "OPTIONS")

	return router
}
