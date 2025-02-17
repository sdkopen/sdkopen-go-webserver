package server

import (
	"log"

	"github.com/sdkopen/sdkopen-go-webserver/logging"
	"github.com/sdkopen/sdkopen-go-webserver/observer"
)

var (
	srvRoutes []Route
	srv       Server
)

type Server interface {
	initialize()
	shutdown() error
	injectMiddlewares()
	injectRoutes()
	listenAndServe() error
}

func AddRoutes(routes []Route) {
	srvRoutes = append(srvRoutes, routes...)
}

func ListenAndServe() {
	srv = createChiServer()
	srv.initialize()
	srv.injectMiddlewares()
	srv.injectRoutes()

	observer.Attach(restObserver{})
	logging.Info("Service '%s' running in %d port", "REST-SERVER", 8080)
	log.Fatal(srv.listenAndServe())
}
