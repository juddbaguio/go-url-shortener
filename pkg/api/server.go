package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/juddbaguio/url-shortener/pkg/config"
	"github.com/juddbaguio/url-shortener/pkg/infra"
	"github.com/juddbaguio/url-shortener/pkg/routes"
)

type Server struct {
	Mux   *mux.Router
	Redis infra.RedisService
}

func NewServer(redis infra.RedisService) *Server {
	mux := mux.NewRouter()
	return &Server{
		Mux:   mux,
		Redis: redis,
	}
}

func (s *Server) Start(cfg config.Cfg) error {
	routes.SetUpRoutes(s.Mux, s.Redis)
	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		Handler:      s.Mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serverError := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("server starting at port: %v", cfg.Port)
		serverError <- srv.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		return fmt.Errorf("server error: %v", err.Error())

	case sig := <-shutdown:
		defer fmt.Printf("shutdown complete: %s", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
			return fmt.Errorf("could not stop the server gracefully: %w", err)
		}
	}

	return nil
}
