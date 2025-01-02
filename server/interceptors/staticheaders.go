package interceptors

import (
	"github.com/pnck-projects/imagevault/internal"
	"net/http"
)

type staticHeadersInterceptor struct{}

var StaticHeadersInterceptor internal.Interceptor = staticHeadersInterceptor{}

func (i staticHeadersInterceptor) Before(w http.ResponseWriter, r *http.Request, cfg *internal.InterceptorConfig) internal.Result {
	h := w.Header()

	h.Set("Content-Type", "application/json")
	// Set HSTS
	h.Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
	h.Set("Access-Control-Allow-Origin", "*")
	h.Set("Access-Control-Allow-Headers", "content-type,accept,x-api-key")
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("X-Frame-Options", "Deny")

	return internal.NotDone()
}
