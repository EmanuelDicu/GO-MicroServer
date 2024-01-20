package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tema-sprc-go/dto"
	"tema-sprc-go/service"
	"tema-sprc-go/utils"

	"github.com/gorilla/mux"
)

func CreateCity(w http.ResponseWriter, r *http.Request) {
	var newCity dto.CreateCity
	json.NewDecoder(r.Body).Decode(&newCity)

	if newCity.NumeOras == "" || newCity.Longitudine == 0 || newCity.Latitudine == 0 {
		utils.Return(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	city, statusCode := service.NewCityOps(r.Context()).CreateCity(newCity)

	utils.Return(w, statusCode, city)
}

func GetAllCities(w http.ResponseWriter, r *http.Request) {
	cities, statusCode := service.NewCityOps(r.Context()).GetAllCities()

	utils.Return(w, statusCode, cities)
}

func GetAllCitiesByCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTara, err := strconv.Atoi(vars["id_tara"])
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err)
		return
	}

	cities, statusCode := service.NewCityOps(r.Context()).GetAllCitiesByCountry(idTara)
	if err != nil {
		utils.Return(w, statusCode, err)
		return
	}

	utils.Return(w, statusCode, cities)
}

func UpdateCity(w http.ResponseWriter, r *http.Request) {
	var newCityData dto.GetCities
	err := json.NewDecoder(r.Body).Decode(&newCityData)
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err)
		return
	}
	newCityData.ID = id

	statusCode := service.NewCityOps(r.Context()).UpdateCity(newCityData)
	if err != nil {
		utils.Return(w, http.StatusInternalServerError, err)
		return
	}

	utils.Return(w, statusCode, nil)
}

func DeleteCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err)
		return
	}

	statusCode := service.NewCityOps(r.Context()).DeleteCity(id)
	if err != nil {
		utils.Return(w, statusCode, err)
		return
	}

	utils.Return(w, statusCode, nil)
}
