package interceptors

import (
	"github.com/pnck-projects/imagevault/internal"
	"net/http"
	"regexp"
)

type paramsInterceptor struct{}

var ParamsInterceptor internal.Interceptor = paramsInterceptor{}

func (i paramsInterceptor) Before(w http.ResponseWriter, r *http.Request, cfg *internal.InterceptorConfig) internal.Result {
	regex, _ := regexp.Compile(cfg.GetRoute().GetPattern() + "$")
	results := regex.FindStringSubmatch(r.URL.Path)
	var params []string
	if len(results) > 1 {
		for _, result := range results[1:] {
			params = append(params, result)
		}
	}
	cfg.GetHandlerConfig().SetRouteParams(params)
	return internal.NotDone()
}
