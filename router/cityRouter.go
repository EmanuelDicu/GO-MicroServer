package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"tema-sprc-go/controller"
)

func registerCityRouter(r *mux.Router) {
	cityRouter := r.PathPrefix("/api/cities").Subrouter()

	// Create city
	cityRouter.HandleFunc("", controller.CreateCity).Methods(http.MethodPost)

	// Get all cities
	cityRouter.HandleFunc("", controller.GetAllCities).Methods(http.MethodGet)

	// Get all cities by country
	cityRouter.HandleFunc("/country/{id_tara:[0-9]+}", controller.GetAllCitiesByCountry).Methods(http.MethodGet)

	// Update city by ID
	cityRouter.HandleFunc("/{id:[0-9]+}", controller.UpdateCity).Methods(http.MethodPut)

	// Delete city by ID
	cityRouter.HandleFunc("/{id:[0-9]+}", controller.DeleteCity).Methods(http.MethodDelete)
}
