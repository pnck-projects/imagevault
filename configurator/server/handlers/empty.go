package handlers

import (
	"github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
)

type empty struct {
}

var EmptyHandler internal.Handler = &empty{}

func (h *empty) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
