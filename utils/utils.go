package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, value any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(value)
}

func WriteJSON(w http.ResponseWriter, value any, status int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(value)
}

func WriteError(w http.ResponseWriter, err error, status int) {
	WriteJSON(w, map[string]string{"error": err.Error()}, status)
}
