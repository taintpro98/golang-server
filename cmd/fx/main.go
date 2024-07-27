package main

import (
	"context"
	"errors"
	"fmt"
	"golang-server/config"
	"golang-server/middleware"
	fx_business "golang-server/module/fx/business"
	"golang-server/module/fx/repository"
	fx_transport "golang-server/module/fx/transport"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"golang-server/pkg/logger"
	"golang-server/route"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

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

func startHTTPServer(lc fx.Lifecycle, cnf config.Config, engine *gin.Engine) {
	server := http.Server{
		Addr:    cnf.AppInfo.ApiPort,
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info(ctx, fmt.Sprintf("Running API on port %s...", cnf.AppInfo.ApiPort))
				err := server.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error(ctx, err, "Run app error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func main() {
	// cnf := config.Init()
	// ctx := context.Background()
	app := fx.New(
		fx.Provide(
			config.Init,
			// func() context.Context {
			// 	return ctx
			// },
		),
		ConnectionModule,
		fx.Provide(fx_transport.NewTransport),
		fx.Provide(NewGinEngine),
		repository.RepositoryModule,
		fx_business.BusinessModule,
		fx.Invoke(startHTTPServer),
	)

	app.Run()
}
