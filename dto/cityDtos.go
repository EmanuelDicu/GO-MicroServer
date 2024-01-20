package dto

type CreateCity struct {
	NumeOras    string  `json:"nume"`
	Latitudine  float64 `json:"lat"`
	Longitudine float64 `json:"lon"`
	IDTara      int     `json:"idTara"`
}
type GetCities struct {
	ID          int     `json:"id"`
	NumeOras    string  `json:"nume"`
	Latitudine  float64 `json:"lat"`
	Longitudine float64 `json:"lon"`
	IDTara      int     `json:"idTara"`
}

type Response struct {
	Id int `json:"id"`
}
