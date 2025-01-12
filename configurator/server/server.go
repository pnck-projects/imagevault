package server

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	handlers2 "github.com/pnck-projects/imagevault/configurator/server/handlers"
)

const (
	objectIdRegex       = `([a-z0-9]{24})`
	objectIdOrNameRegex = `([a-zA-Z0-9\s-_]+)`
)

func Load(mux *CustomMux) {

	// Documentation
	mux.AddRoute(internal2.NewRoute("/docs", "GET", internal2.Public, handlers2.SwaggerRedirectHandler))
	mux.AddRoute(internal2.NewRoute("/docs/swagger.json", "GET", internal2.Public, handlers2.SwaggerJsonHandler))
	mux.AddRoute(internal2.NewRoute("/docs/.*", "GET", internal2.Public, handlers2.SwaggerBaseHandler))

	// Authenticated endpoints
	//mux.AddRoute(vault.NewRoute("/token", "GET", vault.EnvironmentAdmin, handlers.TokenListHandler))

	// Public Endpoints
	mux.AddRoute(internal2.NewRoute("/", "GET", internal2.Public, handlers2.IndexHandler))
	mux.AddRoute(internal2.NewRoute("/version", "GET", internal2.Public, handlers2.VersionHandler))

	// Health checks for application
	mux.AddRoute(internal2.NewRoute("/ready", "GET", internal2.Public, handlers2.ReadinessHandler))
	mux.AddRoute(internal2.NewRoute("/live", "GET", internal2.Public, handlers2.LivenessHandler))

}
