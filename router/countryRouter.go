package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"tema-sprc-go/controller"
)

func registerCountryRouter(r *mux.Router) {

	countryRouter := r.PathPrefix("/api/countries").Subrouter()

	// Create country
	countryRouter.HandleFunc("", controller.CreateCountry).Methods(http.MethodPost)

	// Get all countries
	countryRouter.HandleFunc("", controller.GetCountries).Methods(http.MethodGet)

	// Update country by ID
	countryRouter.HandleFunc("/{id:[0-9]+}", controller.UpdateCountry).Methods(http.MethodPut)

	// Delete country by ID
	countryRouter.HandleFunc("/{id:[0-9]+}", controller.DeleteCountry).Methods(http.MethodDelete)
}
