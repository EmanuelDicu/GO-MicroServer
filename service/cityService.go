// city_service.go
package service

import (
	"context"
	"net/http"
	"tema-sprc-go/config"
	"tema-sprc-go/dto"
	"tema-sprc-go/ent"
	"tema-sprc-go/ent/city"
	"tema-sprc-go/ent/country"
)

type CityOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewCityOps(ctx context.Context) *CityOps {
	return &CityOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (c *CityOps) GetAllCities() ([]dto.GetCities, int) {
	cities, _ := c.client.City.Query().All(c.ctx)

	var citiesDto []dto.GetCities
	for _, ci := range cities {
		citiesDto = append(citiesDto, dto.GetCities{
			ID:          ci.ID,
			NumeOras:    ci.NumeOras,
			Latitudine:  ci.Latitudine,
			Longitudine: ci.Longitudine,
			IDTara:      ci.IDTara,
		})
	}

	return citiesDto, http.StatusOK
}

func (c *CityOps) GetAllCitiesByCountry(idTara int) ([]*dto.GetCities, int) {
	cities, _ := c.client.City.Query().
		Where(
			city.IDTara(idTara),
		).
		All(c.ctx)

	var citiesDto []*dto.GetCities
	for _, ci := range cities {
		citiesDto = append(citiesDto, &dto.GetCities{
			ID:          ci.ID,
			NumeOras:    ci.NumeOras,
			Latitudine:  ci.Latitudine,
			Longitudine: ci.Longitudine,
			IDTara:      ci.IDTara,
		})
	}

	return citiesDto, http.StatusOK
}

func (c *CityOps) CreateCity(newCity dto.CreateCity) (*dto.Response, int) {

	cityExists, _ := c.client.City.Query().
		Where(city.NumeOrasEQ(newCity.NumeOras)).
		Exist(c.ctx)

	if cityExists {
		return nil, http.StatusConflict
	}

	countryExists, _ := c.client.Country.Query().
		Where(country.IDEQ(newCity.IDTara)).
		Exist(c.ctx)

	if !countryExists {
		return nil, http.StatusNotFound
	}

	newCreatedCity, _ := c.client.City.Create().
		SetNumeOras(newCity.NumeOras).
		SetLatitudine(newCity.Latitudine).
		SetLongitudine(newCity.Longitudine).
		SetIDTara(newCity.IDTara).
		Save(c.ctx)

	resp := dto.Response{
		Id: newCreatedCity.ID,
	}

	return &resp, http.StatusCreated
}

func (c *CityOps) UpdateCity(newCity dto.GetCities) int {
	cityExists, _ := c.client.City.Query().
		Where(city.ID(newCity.ID)).
		Exist(c.ctx)

	if !cityExists {
		return http.StatusNotFound
	}

	countryExists, _ := c.client.Country.Query().
		Where(country.IDEQ(newCity.IDTara)).
		Exist(c.ctx)

	if !countryExists {
		return http.StatusNotFound
	}

	//check if there is a city with the same name in the same country
	cityExists, _ = c.client.City.Query().
		Where(city.NumeOrasEQ(newCity.NumeOras), city.IDTaraEQ(newCity.IDTara)).
		Exist(c.ctx)

	if cityExists {
		return http.StatusConflict
	}

	c.client.City.UpdateOneID(newCity.ID).
		SetNumeOras(newCity.NumeOras).
		SetLatitudine(newCity.Latitudine).
		SetLongitudine(newCity.Longitudine).
		SetIDTara(newCity.IDTara).
		Save(c.ctx)

	return http.StatusOK
}

func (c *CityOps) DeleteCity(id int) int {
	cityExists, _ := c.client.City.Query().
		Where(city.ID(id)).
		Exist(c.ctx)

	if !cityExists {
		return http.StatusNotFound
	}

	c.client.City.
		DeleteOneID(id).
		Exec(c.ctx)

	return http.StatusOK
}
