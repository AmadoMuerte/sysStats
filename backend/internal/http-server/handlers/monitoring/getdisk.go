package monitoring

import (
	"net/http"

	"github.com/AmadoMuerte/sysStats/internal/lib/response"
	"github.com/go-chi/render"
	"github.com/shirou/gopsutil/v4/disk"
)

func (m *MonitoringHandler) GetDisk(w http.ResponseWriter, r *http.Request) {
	d, err := disk.Usage("/")
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	render.JSON(w, r, d)
}
