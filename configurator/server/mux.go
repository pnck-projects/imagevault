package server

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	interceptors2 "github.com/pnck-projects/imagevault/configurator/server/interceptors"
	"github.com/pnck-projects/imagevault/database"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type CustomMux struct {
	defaultMux   *http.ServeMux
	routes       []internal2.Route
	interceptors []internal2.Interceptor
	db           *database.MongoDB
}

func NewCustomMux(db *database.MongoDB) CustomMux {
	return CustomMux{db: db}
}

// Ensure http.Handler interface is satisfied
func (m *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.defaultMux == nil {
		m.init()
	}

	cfg := internal2.NewHandlerConfig(m.db)

	if r.Method == "OPTIONS" {
		interceptors2.StaticHeadersInterceptor.Before(w, r, nil)
		w.WriteHeader(http.StatusOK)
		return
	}

	for _, route := range m.routes {

		if r.Method == route.GetMethod() {
			requestedPath := r.URL.Path
			// Strip last slash to match list endpoints
			if len(requestedPath) > 1 && strings.HasSuffix(requestedPath, "/") {
				requestedPath = requestedPath[:len(requestedPath)-1]
			}
			regex, _ := regexp.Compile(route.GetPattern() + "$")

			if regex.MatchString(requestedPath) {
				for _, i := range m.interceptors {
					res := i.Before(w, r, internal2.NewInterceptorConfig(m.db, &route, &cfg))
					if res.Done {
						return
					}
				}
				if requestedPath != "/ready" && requestedPath != "/live" {
					log.Println(route.GetMethod() + " " + route.GetPattern())
				}

				route.Handle(cfg).ServeHTTP(w, r)
				return
			}
		}
	}

	http.NotFound(w, r)
}

func (m *CustomMux) AddRoute(route internal2.Route) {
	m.routes = append(m.routes, route)
}

func (m *CustomMux) init() {
	m.defaultMux = http.NewServeMux()

	m.interceptors = append(m.interceptors, interceptors2.StaticHeadersInterceptor)
	m.interceptors = append(m.interceptors, interceptors2.AuthInterceptor)
	m.interceptors = append(m.interceptors, interceptors2.ParamsInterceptor)

}
