package handlers

import (
	"encoding/json"
	"github.com/pnck-projects/imagevault/configurator/server/responses"
	"net/http"
)

func ThrowError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(responses.NewError(code, message))
	w.WriteHeader(code)
	w.Write(response)
}
