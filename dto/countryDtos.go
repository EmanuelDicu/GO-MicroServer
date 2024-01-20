// Assuming this file is in the "Weather" package
package dto

// CreateCountry is a DTO for creating a country
type CreateCountry struct {
	Nume string  `json:"nume"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

// GetCountries is a DTO for getting countries
type GetCountries struct {
	Id   int     `json:"id"`
	Nume string  `json:"nume"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

// UpdateCountry is a DTO for updating countries
type UpdateCountry struct {
	Id   int     `json:"id"`
	Nume string  `json:"nume,omitempty"`
	Lat  float64 `json:"lat,omitempty"`
	Lon  float64 `json:"lon,omitempty"`
}
