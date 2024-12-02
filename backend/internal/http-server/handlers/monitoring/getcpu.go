package monitoring

import (
	"net/http"

	"github.com/AmadoMuerte/sysStats/internal/lib/response"
	"github.com/go-chi/render"
	"github.com/shirou/gopsutil/v4/cpu"
)

type ProcResponse struct {
}

func (m *MonitoringHandler) GetCPU(w http.ResponseWriter, r *http.Request) {
	c, err := cpu.Info()
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	render.JSON(w, r, c)
}
