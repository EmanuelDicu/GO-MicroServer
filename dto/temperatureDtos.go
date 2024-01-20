package dto

import "time"

// CreateTemperature is a DTO for creating temperatures
type CreateTemperature struct {
	IDOras  int     `json:"idOras"`
	Valoare float64 `json:"valoare"`
}

// GetTemperatures is a DTO for retrieving temperatures
type GetTemperatures struct {
	ID        int       `json:"id"`
	Valoare   float64   `json:"valoare"`
	Timestamp time.Time `json:"timestamp"`
}

// UpdateTemperature is a DTO for updating temperatures
type UpdateTemperature struct {
	ID      int     `json:"id"`
	IDOras  int     `json:"idOras"`
	Valoare float64 `json:"valoare"`
}
