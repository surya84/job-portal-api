package main

import (
	//"job-portal/internal/handlers"
	"context"
	"fmt"
	"job-portal/config"
	"job-portal/internal/auth"
	"job-portal/internal/database"
	"job-portal/internal/handlers"
	rediscache "job-portal/internal/redisCache"
	"job-portal/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("hello this is our app")
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}

}
func startApp() error {
	config.Init()
	cfg := config.GetConfig()

	// =========================================================================
	// Initialize authentication support
	log.Info().Msg("main : Started : Initializing authentication support")
	privatePEM := cfg.Keys.PrivateKey
	// if err != nil {
	// 	return fmt.Errorf("reading auth private key %w", err)
	// }
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	publicPEM := cfg.Keys.PublicKey
	// if err != nil {
	// 	return fmt.Errorf("reading auth public key %w", err)
	// }

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicPEM))
	if err != nil {
		return fmt.Errorf("parsing auth public key %w", err)
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	// =========================================================================
	// Start Database
	log.Info().Msg("main : Started : Initializing db support")
	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w ", err)
	}
	////Initializing redis conn
	rdb := rediscache.RedisClient()
	redisLayer := rediscache.NewredisConnection(rdb)

	// =========================================================================
	//Initialize Conn layer support
	ms, err := repository.NewRepo(db)
	if err != nil {
		return err
	}

	api := http.Server{
		Addr: fmt.Sprintf(":%s", cfg.AppConfig.Port),

		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTime) * time.Second,
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTime) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.Idle_Time) * time.Second,
		Handler:      handlers.API(a, ms, redisLayer),
	}

	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("port", api.Addr).Msg("main: API listening")
		serverErrors <- api.ListenAndServe()
	}()
	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			//Close immediately closes all active net.Listeners
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil

}
