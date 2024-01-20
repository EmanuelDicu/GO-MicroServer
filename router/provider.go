package router

import "github.com/gorilla/mux"

// register all available route
func RegisterRouter(r *mux.Router) {
	registerCountryRouter(r)
	registerCityRouter(r)
	registerTemperatureRouter(r)
}
