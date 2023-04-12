package app

import (
	"llg_backend/config"
	"llg_backend/internal/entity"
	"llg_backend/internal/pkg/postgres"
	"llg_backend/internal/presentation/controller/http"
	"llg_backend/internal/token"
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

	db, err := postgres.New(&cfg.Postgres)
	if err != nil {
		zapLogger.Fatalw("cannot connect to postgres", "err", err)
	}

	db.AutoMigrate(
		&entity.Admin{},
		&entity.User{},
		&entity.SignInHistory{},
		&entity.Item{}, &entity.Door{},
		&entity.Rule{}, &entity.World{},
		&entity.MapConfiguration{},
		&entity.MapConfigurationItem{},
		&entity.MapConfigurationRule{},
		&entity.MapConfigurationDoor{},
		&entity.MapConfigurationForPlayer{},
		&entity.GameSession{},
		&entity.SubmitHistory{},
		&entity.StateValue{},
		&entity.SubmitHistoryRule{},
		&entity.CommandNode{},
		&entity.CommandEdge{},
	)

	tokenMaker, err := token.NewJWTMaker(cfg.JWT.SecretKey)
	if err != nil {
		zapLogger.Fatalw("cannot create token maker", "err", err)
	}

	http.NewRouter(handler, cfg, db, tokenMaker)

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
