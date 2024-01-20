package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tema-sprc-go/dto"
	"tema-sprc-go/service"
	"tema-sprc-go/utils"
)

func CreateCountry(w http.ResponseWriter, r *http.Request) {
	var newCountry dto.CreateCountry
	json.NewDecoder(r.Body).Decode(&newCountry)

	if newCountry.Nume == "" || newCountry.Lon == 0 || newCountry.Lat == 0 {
		utils.Return(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	country, statusCode := service.NewCountryOps(r.Context()).CreateCountry(newCountry)

	utils.Return(w, statusCode, country)
}

func GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, statusCode := service.NewCountryOps(r.Context()).GetCountries()

	utils.Return(w, statusCode, countries)
}

func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	var newCountryData dto.UpdateCountry
	json.NewDecoder(r.Body).Decode(&newCountryData)

	if newCountryData.Nume == "" || newCountryData.Lon == 0 || newCountryData.Lat == 0 {
		utils.Return(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	statusCode := service.NewCountryOps(r.Context()).UpdateCountry(id, newCountryData)

	utils.Return(w, statusCode, nil)
}

func DeleteCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	statusCode, err := service.NewCountryOps(r.Context()).DeleteCountry(id)

	utils.Return(w, statusCode, err)
}
