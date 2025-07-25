package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wenyinh/go-wire-app/api/v1/user"
	"github.com/wenyinh/go-wire-app/pkg/config"
	"github.com/wenyinh/go-wire-app/pkg/logger"
	"github.com/wenyinh/go-wire-app/pkg/middleware/handler"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	Config         *config.AppConfiguration
	UserController *user.Controller
}

func New(config *config.AppConfiguration, userController *user.Controller) *App {
	return &App{
		Config:         config,
		UserController: userController,
	}
}

func (a *App) Serve(rootCtx context.Context) {
	r := gin.New()
	logger.InitLogger()
	if _, err := a.SetupRouter(rootCtx, r); err != nil {
		zap.L().Error("failed to setup router")
		os.Exit(1)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(a.Config.AppConfig.Port),
		Handler: r,
		BaseContext: func(net.Listener) context.Context {
			return rootCtx
		},
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	stopCtx, stop := signal.NotifyContext(rootCtx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-stopCtx.Done()
	zap.L().Info("shutting down server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	closeCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.Shutdown(closeCtx); err != nil {
			zap.L().Error("failed to shutdown server")
		}
	}()
	wg.Wait()
	zap.L().Info("server exiting")
}

func (a *App) SetupRouter(_ context.Context, router *gin.Engine) (*gin.Engine, error) {
	log := zap.L().With(zap.String("AppName", a.Config.AppConfig.AppName))
	log.Info("Setting up Router For App")

	r := router.Group("")
	// Config middleware
	r.Use(handler.GinLogger())
	r.Use(handler.GinRecovery(true))

	r.GET(HealthStatusUri, GetHealthStatus)
	r.HEAD(HealthStatusUri, GetHealthStatus)

	apiV1 := r.Group("/api/v1")
	a.UserController.RegisterRoutes(apiV1)

	return router, nil
}
