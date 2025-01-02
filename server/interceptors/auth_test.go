package interceptors

import (
	"github.com/pnck-projects/imagevault/internal"
	"net/http"
	"testing"
)

func Test_authInterceptor_Before(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	testRoute := internal.NewRoute("/", "GET", internal.EnvironmentAdmin, nil)
	handlerConfig := internal.HandlerConfig{}
	interceptorConfig := internal.NewInterceptorConfig(nil, &testRoute, &handlerConfig)
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
