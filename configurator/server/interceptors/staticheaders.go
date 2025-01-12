package interceptors

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
)

type staticHeadersInterceptor struct{}

var StaticHeadersInterceptor internal2.Interceptor = staticHeadersInterceptor{}

func (i staticHeadersInterceptor) Before(w http.ResponseWriter, r *http.Request, cfg *internal2.InterceptorConfig) internal2.Result {
	h := w.Header()

	h.Set("Content-Type", "application/json")
	// Set HSTS
	h.Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Headers", "content-type,accept,x-api-key")
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("X-Frame-Options", "Deny")

	return internal2.NotDone()
}
