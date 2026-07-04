// Package main is the Travel-Assistant API Gateway (Hertz) entry point
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/fuyx-goai/Travel-Assistant/internal/gateway/router"
	"github.com/fuyx-goai/Travel-Assistant/pkg/config"
	"github.com/fuyx-goai/Travel-Assistant/pkg/logger"
)

var (
	appConfig *config.Config
	appLogger *zap.Logger
	h         *server.Hertz
)

func init() {
	configPath := flag.String("config", "configs", "Config directory path")
	env := flag.String("env", "", "Runtime environment: dev, test, prod")
	flag.Parse()

	var err error
	appConfig, err = config.Load(*configPath, *env)
	if err != nil {
		fmt.Printf("[FATAL] Config load failed: %v\n", err)
		os.Exit(1)
	}

	loggerCfg := logger.LoggerConfig{
		Level:        appConfig.Logger.Level,
		Format:       appConfig.Logger.Format,
		MaxSizeMB:    appConfig.Logger.MaxSizeMB,
		MaxBackups:   appConfig.Logger.MaxBackups,
		MaxAgeDays:   appConfig.Logger.MaxAgeDays,
		Compress:     appConfig.Logger.Compress,
		OutputPaths:  appConfig.Logger.OutputPaths,
		ErrorOutput:  appConfig.Logger.ErrorOutput,
	}

	if err := logger.Init(loggerCfg); err != nil {
		fmt.Printf("[FATAL] Logger init failed: %v\n", err)
		os.Exit(1)
	}
	appLogger = logger.Get()

	_ = initDatabase()
	_ = initRedis()
}

func main() {
	opts := []server.Option{
		server.WithHostPorts(fmt.Sprintf("%s:%d", appConfig.Server.Host, appConfig.Server.Port)),
		server.WithReadTimeout(time.Duration(appConfig.Server.ReadTimeout) * time.Second),
		server.WithWriteTimeout(time.Duration(appConfig.Server.WriteTimeout) * time.Second),
	}

	h = server.Default(opts...)
	registerMiddleware(h)
	router.RegisterRoutes(h)

	go func() {
		if err := h.Run(); err != nil {
			appLogger.Fatal("API gateway startup failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	h.Shutdown(ctx)
}

func registerMiddleware(h *server.Hertz) {
	h.Use(middleware.Recovery(appLogger))
	h.Use(middleware.RequestID())
	h.Use(middleware.CORS())
	h.Use(middleware.Logger(appLogger))
}

func initDatabase() error { return nil }
func initRedis() error { return nil }