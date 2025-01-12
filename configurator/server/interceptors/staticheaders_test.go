package interceptors

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
	"testing"
)

func Test_staticHeadersInterceptor_Before(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	testRoute := internal2.NewRoute("/", "GET", internal2.Public, nil)
	handlerConfig := internal2.HandlerConfig{}
	interceptorConfig := internal2.NewInterceptorConfig(nil, &testRoute, &handlerConfig)
	responseWriter := MockResponseWriter{}
	staticHeadersInterceptor{}.Before(&responseWriter, request, interceptorConfig)

	headersThatShouldExist := []string{"Access-Control-Allow-Methods", "Content-Type", "X-Content-Type-Options", "X-Frame-Options"}

	for _, v := range headersThatShouldExist {
		if responseWriter.Headers.Get(v) == "" {
			t.Fatalf("staticHeadersInterceptor should set header %s", v)
		}
	}

}
