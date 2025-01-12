package interceptors

import (
	internal2 "github.com/pnck-projects/imagevault/configurator/internal"
	"net/http"
)

type authInterceptor struct {
}

const Header = "X-API-KEY"
const UnauthorizedText = "Unauthorized"

var AuthInterceptor internal2.Interceptor = authInterceptor{}

func (i authInterceptor) Before(w http.ResponseWriter, r *http.Request, cfg *internal2.InterceptorConfig) internal2.Result {
	/*if cfg.GetMinimalAuth() > internal.Public {
		header := r.Header.Get(Header)
		if header == "" {
			handlers.ThrowError(w, http.StatusUnauthorized, UnauthorizedText)
			return internal.Done()
		}

		token, err := i.lookup(header, cfg)
		if err != nil {
			log.Println("Invalid token: " + err.Error())
			handlers.ThrowError(w, http.StatusUnauthorized, UnauthorizedText)
			return internal.Done()
		}
		if token.TokenType < cfg.GetMinimalAuth() {
			log.Println("Unauthorized")
			handlers.ThrowError(w, http.StatusUnauthorized, UnauthorizedText)
			return internal.Done()
		}
	}*/

	return internal2.NotDone()
}
