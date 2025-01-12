package main

import (
	"errors"
	"fmt"
	"github.com/pnck-projects/imagevault/authenticator/internal"
	"log"
	"net/http"
	"os"
	"time"
)

// https://medium.com/@adigunhammedolalekan/creating-docker-registry-token-authentication-server-with-go-1ce3aa030c17
// openssl req -x509 -nodes -new -sha256 -days 3650 -newkey rsa:2048 -keyout RootCA.key -out RootCA.pem -subj "/C=NL/CN=localhost"
// openssl x509 -outform pem -in RootCA.pem -out RootCA.crt

func main() {
	crt, key := "/home/gijsm/git/registry_auth/certs/RootCA.crt", "/home/gijsm/git/registry_auth/certs/RootCA.key"
	opt := &internal.Option{
		Certfile:        "/home/gijsm/git/registry_auth/certs/RootCA.crt",
		Keyfile:         "/home/gijsm/git/registry_auth/certs/RootCA.key",
		TokenExpiration: time.Now().Add(24 * time.Hour).Unix(),
		TokenIssuer:     "PNCK",
		Authenticator:   &httpAuthenticator{},
		Authorizer:      &registryAuthorizer{},
	}
	srv, err := internal.NewAuthServer(opt)
	if err != nil {
		log.Fatal(err)
	}
	port := "5011"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := ":" + port
	http.Handle("/auth", srv)
	log.Println("Server running at ", addr)
	if err := http.ListenAndServeTLS(addr, crt, key, nil); err != nil {
		log.Fatal(err)
	}
}

type httpAuthenticator struct {
}

func (h *httpAuthenticator) Authenticate(username, password string) error {
	println(password)
	if username != "gijsm" {
		return errors.New("error")
	}
	return nil
}

type registryAuthorizer struct {
}

func (a *registryAuthorizer) Authorize(req *internal.AuthorizationRequest) ([]string, error) {
	println(fmt.Sprintf("%v", req))

	return []string{"pull", "push"}, nil
}
