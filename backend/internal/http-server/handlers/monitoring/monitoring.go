package monitoring

import (
	"net/http"

	"github.com/AmadoMuerte/sysStats/internal/config"
	"github.com/AmadoMuerte/sysStats/internal/db"
)

type MonitoringHandler struct {
	cfg *config.Config
	db  *db.Storage
}

type IMonitoring interface {
	GetMem(w http.ResponseWriter, r *http.Request)
	GetCPU(w http.ResponseWriter, r *http.Request)
	GetDisk(w http.ResponseWriter, r *http.Request)
}

func New(cfg *config.Config, db *db.Storage) IMonitoring {
	return &MonitoringHandler{cfg, db}
}
