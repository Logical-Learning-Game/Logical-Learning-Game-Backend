package app

import (
	"llg_backend/config"
	v1 "llg_backend/internal/controller/http/v1"
	"llg_backend/internal/entity/sqlc_generated"
	"llg_backend/internal/pkg/postgres"
	"llg_backend/internal/repository"
	"llg_backend/internal/service"
	"llg_backend/pkg/httpserver"
	"llg_backend/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot initialize zap development logger: %v", err.Error())
	}
	defer l.Sync()

	sugar := l.Sugar()
	zapLogger := logger.NewZapLogger(sugar)

	globalLogger := l.Sugar()
	zapGlobalLogger := logger.NewZapLogger(globalLogger)
	logger.SetGlobalLogger(zapGlobalLogger)

	handler := gin.New()

	conn, err := postgres.New(cfg.Postgres.URI)
	if err != nil {
		zapLogger.Fatalw("cannot connect to postgres, exiting program", "err", err)
	}

	postgresQuery := sqlc_generated.New(conn)
	playerRepo := repository.NewPlayerRepository(postgresQuery)
	playerService := service.NewPlayerService(playerRepo)
	playerServiceWithLog := service.NewPlayerServiceWithLog(playerService, zapLogger)

	v1.NewRouter(handler, cfg, playerServiceWithLog)

	httpServer := httpserver.NewServer(handler, httpserver.Port(cfg.HTTP.Port))
	httpServer.Start()

	// Waiting Signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		zapLogger.Infow("interrupt occur", "signal", s)
	case err := <-httpServer.Notify():
		zapLogger.Errorw("http server run", "err", err)
	}

	// Shutdown
	if err := httpServer.Shutdown(); err != nil {
		zapLogger.Errorw("http server shutdown", "err", err)
	}
}
