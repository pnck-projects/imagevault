package interceptors

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
	"testing"
)

func Test_authInterceptor_Before(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	testRoute := internal2.NewRoute("/", "GET", internal2.EnvironmentAdmin, nil)
	handlerConfig := internal2.HandlerConfig{}
	interceptorConfig := internal2.NewInterceptorConfig(nil, &testRoute, &handlerConfig)
	responseWriter := MockResponseWriter{}
	authInterceptor{}.Before(&responseWriter, request, interceptorConfig)

	/*	if responseWriter.Code != 401 {
			t.Fatal("AuthInterceptor should throw unauthorized when no token is sent")
		}
		responseWriter.WriteHeader(0)

		request.Header.Add("X-API-Key", "test")
		authInterceptor{}.Before(&responseWriter, request, interceptorConfig)
		if responseWriter.Code != 401 {
			t.Fatal("AuthInterceptor should throw unauthorized when invalid token is sent")
		}*/
}
