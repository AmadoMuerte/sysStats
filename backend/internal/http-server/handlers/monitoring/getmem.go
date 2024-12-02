package monitoring

import (
	"net/http"

	"github.com/AmadoMuerte/sysStats/internal/lib/response"
	"github.com/go-chi/render"
	"github.com/shirou/gopsutil/v4/mem"
)

func (m *MonitoringHandler) GetMem(w http.ResponseWriter, r *http.Request) {
	v, err := mem.VirtualMemory()
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	render.JSON(w, r, v)
}
