package server

import (
	"log/slog"
	"net/http"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/db"
	"github.com/AmadoMuerte/FlickSynergy/internal/http-server/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	cfg *config.Config
	db  *db.Storage
}

func New(cfg *config.Config, db *db.Storage) *Server {
	return &Server{cfg, db}
}

func (s *Server) Start() {
	router := s.createRouter()

	srv := &http.Server{
		Addr:    s.cfg.App.Address + ":" + s.cfg.App.Port,
		Handler: router,
	}

	slog.Info("server started", slog.Group("app",
		slog.String("address", srv.Addr),
		slog.String("mode", s.cfg.App.Mode),
	))

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("failed to start server", slog.String("error", err.Error()))
		panic(err)
	}
}

func (s *Server) createRouter() http.Handler {

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Разрешаем запросы с любых источников, убрать после релиза
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)
	_ = middlewares.AuthMiddleware{Cfg: s.cfg}

	return router
}
