package webserver

import (
	_ "encoding/json"
	"fmt"
	"github.com/go-chi/cors"
	"github.com/sdkopen/sdkopen-go-webbase/server"
	"net/http"
	"sync"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sdkopen/sdkopen-go-webbase/logging"
	"github.com/sdkopen/sdkopen-go-webbase/observer"
)

type ChiWebServer struct {
	engine *chi.Mux
	srv    *http.Server
	wg     *sync.WaitGroup
}

func CreateChiServer() server.Server {
	return &ChiWebServer{}
}

func (s *ChiWebServer) Initialize() {
	s.engine = chi.NewRouter()
	s.wg = observer.GetWaitGroup()
}

func (s *ChiWebServer) Shutdown() error {
	return s.srv.Close()
}

func (s *ChiWebServer) InjectMiddlewares() {
	s.engine.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // tempo em segundos que a resposta de preflight pode ser cacheada
	}))
	s.engine.Use(middleware.Recoverer)
}

func (s *ChiWebServer) InjectRoutes() {

	for _, route := range server.SrvRoutes {
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

func (s *ChiWebServer) ListenAndServe() error {
	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: s.engine,
	}
	return s.srv.ListenAndServe()
}
