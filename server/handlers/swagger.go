package handlers

import (
	"github.com/pnck-projects/imagevault/internal"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"net/http"
	"os"
)

type swaggerRedirect struct {
}

var SwaggerRedirectHandler internal.Handler = &swaggerRedirect{}

func (h *swaggerRedirect) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/docs/index.html")
		w.WriteHeader(302)
	}
}

type swaggerBase struct {
}

var SwaggerBaseHandler internal.Handler = &swaggerBase{}

func (h *swaggerBase) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"),
	)
}

type swaggerJson struct {
}

var SwaggerJsonHandler internal.Handler = &swaggerJson{}

func (h *swaggerJson) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		swaggerJson, err := os.ReadFile("docs/swagger.json")
		if err != nil {
			ThrowError(w, 404, err.Error())
		}
		w.Write(swaggerJson)
	}
}
