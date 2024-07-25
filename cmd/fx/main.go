package main

import (
	"context"
	"fmt"
	fx_business "golang-server/module/fx/business"
	"golang-server/module/fx/repository"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"net/http"

	"go.uber.org/fx"
)

// Service interface
type Service interface {
	GetData() string
}

// Service implementation
type serviceImpl struct{}

func (s *serviceImpl) GetData() string {
	return "Hello, World!"
}

// Handler that uses the Service
type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := h.service.GetData()
	fmt.Fprintln(w, data)
}

func NewService() Service {
	return &serviceImpl{}
}

func NewMux(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	return mux
}

func startHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}

var ConnectionModule = fx.Module(
	"connection",
	fx.Provide(
		database.NewPostgresqlDatabase,
		cache.NewRedisClient,
	),
)

func main() {
	app := fx.New(
		ConnectionModule,
		repository.RepositoryModule,
		fx_business.BusinessModule,
		// fx.Invoke(func (authenBiz )  {

		// }),
	)

	app.Run()
}
