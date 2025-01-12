package handlers

import (
	"github.com/heptiolabs/healthcheck"
	"github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
	"time"
)

type liveness struct {
}

var LivenessHandler internal.Handler = &liveness{}

// LivenessHandler godoc
//
//	@Summary		Liveness check
//	@Description	Check if application is healthy
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	nil "OK, Healthy"
//	@Failure		500	{string}	string	"Not healthy"
//	@Router			/live [get]
func (h *liveness) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := healthcheck.NewHandler()
		health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
		health.ReadyEndpoint(w, r)
	}
}

type readiness struct {
}

var ReadinessHandler internal.Handler = &readiness{}

// ReadinessHandler godoc
//
//	@Summary		Readiness check
//	@Description	Check if application is ready
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	nil "OK, Ready"
//	@Failure		500	{string}	string	"Not ready"
//	@Router			/ready [get]
func (h *readiness) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := healthcheck.NewHandler()
		health.AddReadinessCheck(
			"upstream-dep-dns",
			healthcheck.DNSResolveCheck("github.com", 50*time.Millisecond))
		//health.AddReadinessCheck("mongodb-connection", healthcheck.Async(checkMongoDbConnection(cfg.GetDb()), 3000))
		health.ReadyEndpoint(w, r)

	}
}
