package server

import (
	"log/slog"
	"net/http"

	_ "github.com/AmadoMuerte/sysStats/docs"
	"github.com/AmadoMuerte/sysStats/internal/config"
	"github.com/AmadoMuerte/sysStats/internal/db"
	authhandler "github.com/AmadoMuerte/sysStats/internal/http-server/handlers/auth"
	monitoringhandler "github.com/AmadoMuerte/sysStats/internal/http-server/handlers/monitoring"
	_ "github.com/AmadoMuerte/sysStats/internal/http-server/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
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
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	apiHandler := authhandler.New(s.cfg, s.db)
	auth := chi.NewRouter()
	auth.Post("/sign-in", apiHandler.SignIn)
	// auth.Post("/sign-up", apiHandler.SignUp)
	auth.Post("/refresh", apiHandler.Refresh)

	monitoring := chi.NewRouter()
	// mw := middlewares.AuthMiddleware{Cfg: s.cfg}
	// monitoring.Use(mw.New)
	monitoringhandler := monitoringhandler.New(s.cfg, s.db)

	monitoring.Get("/mem", monitoringhandler.GetMem)
	monitoring.Get("/cpu", monitoringhandler.GetCPU)
	monitoring.Get("/disk", monitoringhandler.GetDisk)

	devMode(s.cfg.App.Mode, s.cfg.App.Address, s.cfg.App.Port, router)
	router.Mount("/api/v1/login", auth)
	router.Mount("/api/v1/monitoring", monitoring)

	return router
}

func devMode(mode string, addr string, port string, r *chi.Mux) {
	if mode == "dev" {
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://"+addr+":"+port+"/swagger/doc.json"),
			httpSwagger.UIConfig(map[string]string{
				"layout": `"BaseLayout"`, //
			}),
		))
		slog.Info("swagger started", slog.String("address", "http://"+addr+":"+port+"/swagger/"))
	}
}
