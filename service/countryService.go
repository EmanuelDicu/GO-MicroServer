package service

import (
	"context"
	"fmt"
	"net/http"
	"tema-sprc-go/config"
	"tema-sprc-go/dto"
	"tema-sprc-go/ent"
	"tema-sprc-go/ent/country"
)

type CountryOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewCountryOps(ctx context.Context) *CountryOps {
	return &CountryOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (c *CountryOps) GetCountries() ([]*dto.GetCountries, int) {
	countries, _ := c.client.Country.Query().All(c.ctx)

	var countriesDto []*dto.GetCountries
	for _, ci := range countries {
		countriesDto = append(countriesDto, &dto.GetCountries{
			Id:   ci.ID,
			Nume: ci.NumeTara,
			Lat:  ci.Latitudine,
			Lon:  ci.Longitudine,
		})
	}

	return countriesDto, http.StatusOK
}

func (c *CountryOps) CreateCountry(newCountry dto.CreateCountry) (*dto.Response, int) {
	// Check if a country with the same name already exists
	countryWithSameName, _ := c.client.Country.Query().
		Where(country.NumeTaraEQ(newCountry.Nume)).
		Exist(c.ctx)

	if countryWithSameName {
		return nil, http.StatusConflict
	}

	// No country with the same name, proceed with creating a new one
	newCreatedCountry, _ := c.client.Country.Create().
		SetNumeTara(newCountry.Nume).
		SetLatitudine(newCountry.Lat).
		SetLongitudine(newCountry.Lon).
		Save(c.ctx)

	resp := dto.Response{
		Id: newCreatedCountry.ID,
	}

	return &resp, http.StatusCreated
}

func (c *CountryOps) UpdateCountry(id int, input dto.UpdateCountry) int {

	// Check if a country with the specified ID exists
	countryExists, _ := c.client.Country.Query().
		Where(country.IDEQ(id)).
		Exist(c.ctx)

	if !countryExists {
		return http.StatusNotFound
	}

	// Check if a country with the same name already exists (except the one with the specified ID)
	countryWithSameName, _ := c.client.Country.Query().
		Where(
			country.NumeTaraEQ(input.Nume),
			country.IDNEQ(id),
		).
		Exist(c.ctx)

	if countryWithSameName {
		return http.StatusConflict
	}

	c.client.Country.UpdateOneID(input.Id).
		SetNumeTara(input.Nume).
		SetLatitudine(input.Lat).
		SetLongitudine(input.Lon).
		Save(c.ctx)

	return http.StatusOK
}

func (c *CountryOps) DeleteCountry(id int) (int, error) {

	countryExists, _ := c.client.Country.Query().
		Where(country.IDEQ(id)).
		Exist(c.ctx)

	if !countryExists {
		return http.StatusNotFound, fmt.Errorf("country with ID %d does not exist", id)
	}

	c.client.Country.
		DeleteOneID(id).
		Exec(c.ctx)

	return http.StatusOK, nil
}
