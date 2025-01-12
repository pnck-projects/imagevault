package interceptors

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
	"regexp"
)

type paramsInterceptor struct{}

var ParamsInterceptor internal2.Interceptor = paramsInterceptor{}

func (i paramsInterceptor) Before(w http.ResponseWriter, r *http.Request, cfg *internal2.InterceptorConfig) internal2.Result {
	regex, _ := regexp.Compile(cfg.GetRoute().GetPattern() + "$")
	results := regex.FindStringSubmatch(r.URL.Path)
	var params []string
	if len(results) > 1 {
		for _, result := range results[1:] {
			params = append(params, result)
		}
	}
	cfg.GetHandlerConfig().SetRouteParams(params)
	return internal2.NotDone()
}
