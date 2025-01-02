package handlers

import (
	"github.com/pnck-projects/imagevault/internal"
	"net/http"
)

type index struct {
}

var IndexHandler internal.Handler = &index{}

func (h *index) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", "https://github.com/pnck-projects/imagevault")
		w.WriteHeader(302)
	}
}
