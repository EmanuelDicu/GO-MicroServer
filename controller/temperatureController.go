package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tema-sprc-go/dto"
	"tema-sprc-go/service"
	"tema-sprc-go/utils"
	"time"
)

func CreateTemperature(w http.ResponseWriter, r *http.Request) {
	var newTemperature dto.CreateTemperature
	json.NewDecoder(r.Body).Decode(&newTemperature)

	if newTemperature.IDOras == 0 || newTemperature.Valoare == 0 {
		utils.Return(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	temperature, statusCode := service.NewTemperatureOps(r.Context()).CreateTemperature(newTemperature)

	utils.Return(w, statusCode, temperature)
}

func GetTemperatures(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	startDate, _ := time.Parse("2006-01-02", queryValues.Get("from"))
	endDate, _ := time.Parse("2006-01-02", queryValues.Get("until"))
	lat, _ := strconv.ParseFloat(queryValues.Get("lat"), 64)
	lon, _ := strconv.ParseFloat(queryValues.Get("lon"), 64)

	temperatures, statusCode := service.NewTemperatureOps(r.Context()).GetTemperatures(lat, lon, startDate, endDate)

	utils.Return(w, statusCode, temperatures)
}

func GetTemperaturesByCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idOras, _ := strconv.Atoi(vars["id_oras"])

	queryValues := r.URL.Query()
	startDate, _ := time.Parse(time.RFC3339, queryValues.Get("from"))
	endDate, _ := time.Parse(time.RFC3339, queryValues.Get("until"))

	temperatures, statusCode := service.NewTemperatureOps(r.Context()).GetTemperaturesByCity(idOras, startDate, endDate)

	utils.Return(w, statusCode, temperatures)
}

func GetTemperaturesByCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTara, _ := strconv.Atoi(vars["id_tara"])

	queryValues := r.URL.Query()

	startDate, _ := time.Parse(time.RFC3339, queryValues.Get("from"))
	endDate, _ := time.Parse(time.RFC3339, queryValues.Get("until"))

	temperatures, statusCode := service.NewTemperatureOps(r.Context()).GetTemperaturesByCountry(idTara, startDate, endDate)

	utils.Return(w, statusCode, temperatures)
}

func UpdateTemperature(w http.ResponseWriter, r *http.Request) {
	var newTemperatureData dto.UpdateTemperature
	err := json.NewDecoder(r.Body).Decode(&newTemperatureData)
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	newTemperatureData.ID = id

	statusCode := service.NewTemperatureOps(r.Context()).UpdateTemperature(newTemperatureData)

	utils.Return(w, statusCode, nil)
}

func DeleteTemperature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	statusCode := service.NewTemperatureOps(r.Context()).DeleteTemperature(id)

	utils.Return(w, statusCode, nil)
}
