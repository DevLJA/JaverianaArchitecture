package utils

import (
	"CircuitBreaker-LauraJoya/dtos"
	"encoding/json"
	"net/http"
)

func RespondWithBody(w http.ResponseWriter, status int, body interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	response := &dtos.APIResponse{}
	response.Data = body

	if err != nil {
		response.Error = err.Error()
	}

	jsonToReturn, errMarshall := json.Marshal(response)

	if errMarshall != nil {
		http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(jsonToReturn)
}
