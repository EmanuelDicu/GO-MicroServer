package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"tema-sprc-go/controller"
)

func registerTemperatureRouter(r *mux.Router) {
	temperatureRouter := r.PathPrefix("/api/temperatures").Subrouter()

	// Create temperature
	temperatureRouter.HandleFunc("", controller.CreateTemperature).Methods(http.MethodPost)

	// Get all temperatures
	temperatureRouter.HandleFunc("", controller.GetTemperatures).
		Methods(http.MethodGet)

	// Get temperatures by city
	temperatureRouter.HandleFunc("/cities/{id_oras:[0-9]+}", controller.GetTemperaturesByCity).
		Methods(http.MethodGet)

	// Get temperatures by country
	temperatureRouter.HandleFunc("/countries/{id_tara:[0-9]+}", controller.GetTemperaturesByCountry).
		Methods(http.MethodGet)

	// Update temperature by ID
	temperatureRouter.HandleFunc("/{id:[0-9]+}", controller.UpdateTemperature).Methods(http.MethodPut)

	// Delete temperature by ID
	temperatureRouter.HandleFunc("/{id:[0-9]+}", controller.DeleteTemperature).Methods(http.MethodDelete)
}
