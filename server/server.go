package server

import (
	"github.com/pnck-projects/imagevault/internal"
	"github.com/pnck-projects/imagevault/server/handlers"
)

const (
	objectIdRegex       = `([a-z0-9]{24})`
	objectIdOrNameRegex = `([a-zA-Z0-9\s-_]+)`
)

func Load(mux *CustomMux) {

	// Documentation
	mux.AddRoute(internal.NewRoute("/docs", "GET", internal.Public, handlers.SwaggerRedirectHandler))
	mux.AddRoute(internal.NewRoute("/docs/swagger.json", "GET", internal.Public, handlers.SwaggerJsonHandler))
	mux.AddRoute(internal.NewRoute("/docs/.*", "GET", internal.Public, handlers.SwaggerBaseHandler))

	// Authenticated endpoints
	//mux.AddRoute(vault.NewRoute("/token", "GET", vault.EnvironmentAdmin, handlers.TokenListHandler))

	// Public Endpoints
	mux.AddRoute(internal.NewRoute("/", "GET", internal.Public, handlers.IndexHandler))
	mux.AddRoute(internal.NewRoute("/version", "GET", internal.Public, handlers.VersionHandler))

	// Health checks for application
	mux.AddRoute(internal.NewRoute("/ready", "GET", internal.Public, handlers.ReadinessHandler))
	mux.AddRoute(internal.NewRoute("/live", "GET", internal.Public, handlers.LivenessHandler))

}
