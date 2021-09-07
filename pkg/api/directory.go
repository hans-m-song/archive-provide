package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hans-m-song/archive-provide/pkg/middleware"
	"github.com/hans-m-song/archive-provide/pkg/provide"
)

const (
	ApiRoot     = "/api"
	GraphqlRoot = "/graphql"
)

var (
	ApiHealth     = "/health"
	ApiAuthor     = "/author"
	ApiTag        = "/tag"
	ApiPublisher  = "/publisher"
	ApiCollection = "/collection"
)

func NewApiHandler(provider *provide.Provider) http.Handler {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeJson)

	routerApi := router.
		PathPrefix(ApiRoot).Subrouter()

	routerApi.
		PathPrefix(ApiHealth).
		Methods(http.MethodGet).
		HandlerFunc(healthHandler)

	routerApi.
		PathPrefix(ApiAuthor).
		Methods(http.MethodGet).
		HandlerFunc(listNamesHandler(provider, "author"))

	routerApi.
		PathPrefix(ApiTag).
		Methods(http.MethodGet).
		HandlerFunc(listNamesHandler(provider, "tag"))

	routerApi.
		PathPrefix(ApiPublisher).
		Methods(http.MethodGet).
		HandlerFunc(listNamesHandler(provider, "publisher"))

	routerApi.
		PathPrefix(ApiCollection).
		Methods(http.MethodGet).
		HandlerFunc(listNamesHandler(provider, "collection"))

	router.
		Path(GraphqlRoot).
		HandlerFunc(defaultNotFoundHandler)

	return router
}
