package service

import (
	"context"
	"net/http"
	"tema-sprc-go/config"
	"tema-sprc-go/dto"
	"tema-sprc-go/ent"
	"tema-sprc-go/ent/city"
	"tema-sprc-go/ent/country"
	"tema-sprc-go/ent/temperature"
	"time"
)

type TemperatureOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewTemperatureOps(ctx context.Context) *TemperatureOps {
	return &TemperatureOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (t *TemperatureOps) CreateTemperature(newTemperature dto.CreateTemperature) (*dto.Response, int) {

	cityExists, _ := t.client.City.Query().
		Where(city.IDEQ(newTemperature.IDOras)).
		Exist(t.ctx)

	if !cityExists {
		return nil, http.StatusNotFound
	}

	existingTemperature, _ := t.client.Temperature.Query().
		Where(
			temperature.IDOras(newTemperature.IDOras),
			temperature.Valoare(newTemperature.Valoare),
		).
		Exist(t.ctx)

	if existingTemperature {
		return nil, http.StatusConflict
	}

	tmp, err := t.client.Temperature.Create().
		SetValoare(newTemperature.Valoare).
		SetTimestamp(time.Now()).
		SetIDOras(newTemperature.IDOras).
		Save(t.ctx)

	if err != nil {
		return nil, http.StatusInternalServerError
	}

	resp := dto.Response{
		Id: tmp.ID,
	}

	return &resp, http.StatusCreated
}

func (t *TemperatureOps) GetTemperatures(lat float64, lon float64, startDate time.Time, endDate time.Time) ([]*dto.GetTemperatures, int) {

	citiesQuery := t.client.City.Query()

	if lat != 0 {
		citiesQuery = citiesQuery.Where(city.LatitudineEQ(lat))
	}

	if lon != 0 {
		citiesQuery = citiesQuery.Where(city.LongitudineEQ(lon))
	}

	cities, _ := citiesQuery.
		Select(city.FieldID).
		All(t.ctx)

	cityIds := make([]int, len(cities))
	for i, c := range cities {
		cityIds[i] = c.ID
	}

	temperaturesQuery := t.client.Temperature.Query().
		Where(temperature.IDOrasIn(cityIds...))

	if startDate != (time.Time{}) {
		temperaturesQuery = temperaturesQuery.Where(temperature.TimestampGTE(startDate))
	}

	if endDate != (time.Time{}) {
		temperaturesQuery = temperaturesQuery.Where(temperature.TimestampLTE(endDate))
	}

	temperatures, _ := temperaturesQuery.All(t.ctx)

	temperaturesDto := make([]*dto.GetTemperatures, len(temperatures))

	for i, temp := range temperatures {
		temperaturesDto[i] = &dto.GetTemperatures{
			ID:        temp.ID,
			Valoare:   temp.Valoare,
			Timestamp: temp.Timestamp,
		}
	}

	return temperaturesDto, http.StatusOK
}

func (t *TemperatureOps) GetTemperaturesByCity(idOras int, startDate time.Time, endDate time.Time) ([]*dto.GetTemperatures, int) {

	cityExists, _ := t.client.City.Query().
		Where(city.IDEQ(idOras)).
		Exist(t.ctx)

	if !cityExists {
		return nil, http.StatusNotFound
	}

	temperatureQuery := t.client.Temperature.Query().
		Where(temperature.IDOras(idOras))

	if startDate != (time.Time{}) {
		temperatureQuery = temperatureQuery.Where(temperature.TimestampGTE(startDate))
	}

	if endDate != (time.Time{}) {
		temperatureQuery = temperatureQuery.Where(temperature.TimestampLTE(endDate))
	}

	temperatures, _ := temperatureQuery.All(t.ctx)

	temperaturesDto := make([]*dto.GetTemperatures, len(temperatures))
	for i, temp := range temperatures {
		temperaturesDto[i] = &dto.GetTemperatures{
			ID:        temp.ID,
			Valoare:   temp.Valoare,
			Timestamp: temp.Timestamp,
		}
	}

	return temperaturesDto, http.StatusOK
}

func (t *TemperatureOps) GetTemperaturesByCountry(idTara int, startDate time.Time, endDate time.Time) ([]*dto.GetTemperatures, int) {

	countryExists, _ := t.client.Country.Query().
		Where(country.IDEQ(idTara)).
		Exist(t.ctx)

	if !countryExists {
		return nil, http.StatusNotFound
	}

	cities, _ := t.client.City.Query().
		Where(city.IDTara(idTara)).
		All(t.ctx)

	cityIDs := make([]int, len(cities))
	for i, c := range cities {
		cityIDs[i] = c.ID
	}

	temperaturesQuery := t.client.Temperature.Query().
		Where(temperature.IDOrasIn(cityIDs...))

	if startDate != (time.Time{}) {
		temperaturesQuery = temperaturesQuery.Where(temperature.TimestampGTE(startDate))
	}

	if endDate != (time.Time{}) {
		temperaturesQuery = temperaturesQuery.Where(temperature.TimestampLTE(endDate))
	}

	temperatures, _ := temperaturesQuery.All(t.ctx)

	temperaturesDto := make([]*dto.GetTemperatures, len(temperatures))
	for i, temp := range temperatures {
		temperaturesDto[i] = &dto.GetTemperatures{
			ID:        temp.ID,
			Valoare:   temp.Valoare,
			Timestamp: temp.Timestamp,
		}
	}

	return temperaturesDto, http.StatusOK
}

func (t *TemperatureOps) UpdateTemperature(newTemperatureData dto.UpdateTemperature) int {
	tmp, _ := t.client.Temperature.Get(t.ctx, newTemperatureData.ID)
	if tmp == nil {
		return http.StatusNotFound
	}

	if tmp.IDOras == newTemperatureData.IDOras && tmp.Valoare == newTemperatureData.Valoare {
		return http.StatusConflict
	}

	if newTemperatureData.IDOras != 0 {
		tmp.IDOras = newTemperatureData.IDOras
	}

	if newTemperatureData.Valoare != 0 {
		tmp.Valoare = newTemperatureData.Valoare
	}

	_, err := t.client.Temperature.UpdateOneID(tmp.ID).
		SetValoare(tmp.Valoare).
		SetIDOras(tmp.IDOras).
		Save(t.ctx)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (t *TemperatureOps) DeleteTemperature(id int) int {
	tempExists, _ := t.client.Temperature.Query().
		Where(temperature.IDEQ(id)).
		Exist(t.ctx)

	if !tempExists {
		return http.StatusNotFound
	}

	t.client.Temperature.
		DeleteOneID(id).
		Exec(t.ctx)

	return http.StatusOK
}
