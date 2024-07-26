package main

import (
	"context"
	"fmt"
	"golang-server/config"
	"golang-server/middleware"
	fx_business "golang-server/module/fx/business"
	"golang-server/module/fx/repository"
	fx_transport "golang-server/module/fx/transport"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"golang-server/route"
	"net/http"

	"github.com/gin-gonic/gin"
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

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(
		middleware.LogRequestInfo(),
		gin.Recovery(),
	)
	engine.Static("/static", "./static")

	// Route chính hiển thị form upload
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	route.RegisterHealthCheckRoute(engine)
	return engine
}

func main() {
	// cnf := config.Init()
	ctx := context.Background()
	app := fx.New(
		fx.Provide(
			config.Init,
			func() context.Context {
				return ctx
			},
		),
		ConnectionModule,
		fx.Provide(fx_transport.NewTransport),
		fx.Provide(NewGinEngine),
		repository.RepositoryModule,
		fx_business.BusinessModule,
		// fx.Invoke(func (authenBiz )  {

		// }),
	)

	app.Run()
}
