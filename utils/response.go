package utils

import (
	"encoding/json"
	"net/http"
)

func Return(w http.ResponseWriter, statusCode int, response any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
