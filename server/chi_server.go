package server

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sdkopen/sdkopen-go-webserver/logging"
	"github.com/sdkopen/sdkopen-go-webserver/observer"
)

type ChiWebServer struct {
	engine *chi.Mux
	srv    *http.Server
	wg     *sync.WaitGroup
}

func CreateChiServer() Server {
	return &ChiWebServer{}
}

func (s *ChiWebServer) initialize() {
	s.engine = chi.NewRouter()
	s.wg = observer.GetWaitGroup()
}

func (s *ChiWebServer) shutdown() error {
	return s.srv.Close()
}

func (s *ChiWebServer) injectMiddlewares() {
	s.engine.Use(middleware.Recoverer)
}

func (s *ChiWebServer) injectRoutes() {

	for _, route := range srvRoutes {
		routeUri := string(route.Prefix) + route.URI
		fn := route.Function

		s.engine.MethodFunc(route.Method.String(), routeUri, func(w http.ResponseWriter, r *http.Request) {
			s.wg.Add(1)
			defer s.wg.Done()
			webContext := &chiWebContext{writer: w, request: r}

			fn(webContext)
		})

		logging.Info("Registered route [%7s] %s", route.Method, routeUri)
	}
}

func (s *ChiWebServer) listenAndServe() error {
	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: s.engine,
	}
	return s.srv.ListenAndServe()
}
