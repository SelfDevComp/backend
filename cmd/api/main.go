package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/sklyar-vlad/selfDev/database"
	"github.com/sklyar-vlad/selfDev/internal/config"
	"github.com/sklyar-vlad/selfDev/internal/handler"
	authHand "github.com/sklyar-vlad/selfDev/internal/handler/auth"
	habitHand "github.com/sklyar-vlad/selfDev/internal/handler/habit"
	userHand "github.com/sklyar-vlad/selfDev/internal/handler/user"
	authAdapt "github.com/sklyar-vlad/selfDev/internal/integrations/casdoor"
	authRepo "github.com/sklyar-vlad/selfDev/internal/repository/auth"
	habitRepo "github.com/sklyar-vlad/selfDev/internal/repository/habit"
	userRepo "github.com/sklyar-vlad/selfDev/internal/repository/user"
	authSrv "github.com/sklyar-vlad/selfDev/internal/service/auth"
	habitSrv "github.com/sklyar-vlad/selfDev/internal/service/habit"
	userSrv "github.com/sklyar-vlad/selfDev/internal/service/user"
	customLogger "github.com/sklyar-vlad/selfDev/logger"
	"github.com/sklyar-vlad/selfDev/middleware"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed create config: %v", err)
	}

	logger, err := customLogger.NewLogger(cfg.Logger)
	if err != nil {
		log.Fatal("failed create logger", zap.Error(err))
	}

	defer func() {
		_ = logger.Sync()
	}()

	ctx := context.Background()

	pool, err := database.NewPostgres(ctx, cfg.Database)
	if err != nil {
		logger.Fatal("failed connect to the postgres", zap.Error(err))
	}
	defer pool.Close()

	redis, err := database.NewRedis(ctx, cfg.Database)
	if err != nil {
		logger.Fatal("failed connect to the redis", zap.Error(err))
	}
	defer func() {
		_ = redis.Close()
	}()

	userRepository := userRepo.NewRepository(pool, logger)
	habitRepository := habitRepo.NewRepository(pool, logger)
	authRepository := authRepo.NewRepository(pool, redis, logger)

	authAdapter := authAdapt.NewAdapter(cfg.Auth)

	userService := userSrv.NewService(userRepository, logger)
	authService := authSrv.NewService(userService, authAdapter, authRepository, cfg.JWT, logger)
	habitService := habitSrv.NewService(habitRepository, userService, logger)

	authHandler := authHand.NewHandler(authService, logger)
	userHandler := userHand.NewHandler(userService, logger)
	habitHandler := habitHand.NewHandler(habitService, logger)

	rootMux := http.NewServeMux()
	handler.RegisterPublicRoutes(rootMux, authHandler)
	protectedMux := http.NewServeMux()
	handler.RegisterProtectedRoutes(protectedMux, userHandler, habitHandler)
	sessionMiddleware := middleware.NewSessionMiddleware(authRepository)
	rootMux.Handle("/api/", sessionMiddleware.Middleware(protectedMux))
	wrapped := middleware.CORS(rootMux, cfg.Server.Middleware)
	service := &http.Server{
		Addr:         ":8080",
		Handler:      wrapped,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	go func() {
		logger.Info("service started at port 8080.")

		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server failed", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Info("shutdown signal received...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := service.Shutdown(ctxShutdown); err != nil {
		logger.Error("graceful shutdown failed", zap.Error(err))
	} else {
		logger.Info("server stopped gracefully")
	}
}
