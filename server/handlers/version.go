package handlers

import (
	"encoding/json"
	"github.com/pnck-projects/imagevault/internal"
	"github.com/pnck-projects/imagevault/server/models"

	"net/http"
	"os"
	"strings"
)

type version struct {
}

var VersionHandler internal.Handler = &version{}

func (h *version) Handle(cfg internal.HandlerConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		textVersion := "non-release"
		versionFile, err := os.ReadFile("version")
		if err == nil {
			textVersion = strings.TrimSpace(string(versionFile))
		}
		version := models.Version{Version: textVersion, Name: cfg.GetDb().Connection.Name()}
		response, _ := json.Marshal(version)

		w.Write(response)
	}
}
