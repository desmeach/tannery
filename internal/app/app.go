package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	authService "tanneryApp/internal/application/auth"
	httpDelivery "tanneryApp/internal/delivery/http"
	"tanneryApp/internal/delivery/http/v1/auth"
	"tanneryApp/internal/infrastructure/config"
	"tanneryApp/internal/infrastructure/jwt"
	authRepository "tanneryApp/internal/infrastructure/repository/mongo/auth"
	"tanneryApp/internal/lib/logger/slogpretty"
	mongoDatabase "tanneryApp/pkg/database/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(configPath string) error {
	log := setupPrettySlog()
	log.Info("Starting application...")

	log.Info("🍃Initialize application`s config...")
	cfg, err := config.Init(configPath)
	if err != nil {
		return err
	}

	log.Info("🍃Creating MongoDB client...")
	mongoClient, err := mongoDatabase.NewClient(cfg.Mongo.URI, cfg.Mongo.Username, cfg.Mongo.Password, cfg.Mongo.MaxPool, cfg.Mongo.MinPool)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("🍃Creating jwt manager...")
	jwtManager := jwt.NewManager([]byte(cfg.JWT.Secret))

	db := mongoClient.Database(cfg.Mongo.CoreDB)

	authRepo := authRepository.NewRepository(log, db)
	//groupRepo := groupRepository.NewRepository(db)

	authServ := authService.NewUserService(log, authRepo, jwtManager)

	authHandler := auth.NewHandler(authServ)

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	httpDelivery.RegisterRoutes(router, cfg, authHandler)

	log.Info("🍃Starting server...")

	hostUrl := cfg.Http.Host + ":" + cfg.Http.Port
	server := &http.Server{
		Addr:         hostUrl,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Info(fmt.Sprintf("🚀Server started on %s", hostUrl))

	err = server.ListenAndServe()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Info("Gracefully stopped")

	return nil
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
