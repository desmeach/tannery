package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	authServices "tanneryApp/internal/application/auth/services"
	userServices "tanneryApp/internal/application/user/services"
	httpDelivery "tanneryApp/internal/delivery/http"
	"tanneryApp/internal/delivery/http/v1/auth"
	"tanneryApp/internal/delivery/http/v1/user"
	"tanneryApp/internal/infrastructure/config"
	"tanneryApp/internal/infrastructure/jwt"
	"tanneryApp/internal/infrastructure/logger/slogpretty"
	authRepository "tanneryApp/internal/infrastructure/repository/mongo/auth"
	profileRepository "tanneryApp/internal/infrastructure/repository/mongo/profile"
	mongoDatabase "tanneryApp/pkg/database/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() error {
	log := setupPrettySlog()
	log.Info("Starting application...")

	log.Info("🍃Initialize application`s config...")
	cfg := config.MustLoad()

	log.Info("🍃Creating MongoDB client...")
	mongoClient, err := mongoDatabase.NewClient(cfg.Mongo.URI, cfg.Mongo.Username, cfg.Mongo.Password, cfg.Mongo.MaxPool, cfg.Mongo.MinPool)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("🍃Creating jwt manager...")
	jwtManager := jwt.NewManager([]byte(cfg.JWT.Secret))

	db := mongoClient.Database(cfg.Mongo.CoreDB)

	authRepo := authRepository.NewRepository(*db)
	profileRepo := profileRepository.NewRepository(*db)

	authServ := authServices.NewUserService(*log, authRepo, *jwtManager)
	profileServ := userServices.NewProfileService(*log, profileRepo, *jwtManager)

	authHandler := auth.NewHandler(authServ)
	userHandler := user.NewHandler(profileServ)

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	httpDelivery.RegisterRoutes(router, cfg, *jwtManager, authHandler, userHandler)

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

	go func() {
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err.Error())
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Info("Shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Common.GracefulTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Server gracefully stopped")
	}

	log.Info("Cleanup complete")

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
