package internal

import (
	"github.com/pnck-projects/imagevault/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
)

type HandlerConfig struct {
	db                *database.MongoDB
	tokenType         TokenType
	environmentId     bson.ObjectID
	tokenId           bson.ObjectID
	tokenDescription  string
	routeParams       []string
	environmentSecret []byte
}

func NewHandlerConfig(db *database.MongoDB) HandlerConfig {
	return HandlerConfig{
		db: db,
	}
}

type Handler interface {
	Handle(cfg HandlerConfig) http.HandlerFunc
}

func (h *HandlerConfig) GetDb() *database.MongoDB {
	return h.db
}

func (h *HandlerConfig) GetTokenType() TokenType {
	return h.tokenType
}

func (h *HandlerConfig) SetTokenType(t TokenType) {
	h.tokenType = t
}

func (h *HandlerConfig) GetEnvironmentId() bson.ObjectID {
	return h.environmentId
}

func (h *HandlerConfig) SetEnvironmentId(id bson.ObjectID) {
	h.environmentId = id
}

func (h *HandlerConfig) GetTokenDescription() string {
	return h.tokenDescription
}

func (h *HandlerConfig) SetTokenDescription(description string) {
	h.tokenDescription = description
}

func (h *HandlerConfig) GetTokenId() bson.ObjectID {
	return h.tokenId
}

func (h *HandlerConfig) SetTokenId(tokenId bson.ObjectID) {
	h.tokenId = tokenId
}

func (h *HandlerConfig) GetRouteParams() []string {
	return h.routeParams
}

func (h *HandlerConfig) SetRouteParams(routeParams []string) {
	h.routeParams = routeParams
}

func (h *HandlerConfig) GetRouteParam(i int) string {
	if len(h.routeParams) < i+1 {
		return ""
	}
	return h.routeParams[i]
}

func (h *HandlerConfig) GetEnvironmentSecret() []byte {
	return h.environmentSecret
}

func (h *HandlerConfig) SetEnvironmentSecret(secret []byte) {
	h.environmentSecret = secret
}
