package websockethandler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		// Разрешаем доступ только с определенных источников
		return origin == "http://localhost:5173" || origin == "http://127.0.0.1:3000"
	},
}

type WebSocketHandler struct {
}

func New() *WebSocketHandler {
	return &WebSocketHandler{}
}

type Metrics struct {
	MemUsed uint64               `json:"mem"`
	Net     []net.IOCountersStat `json:"net"`
	Cpu     []float64            `json:"cpu"`
	Time    int64                `json:"time"`
}

func getSystemMetrics() (Metrics, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		slog.Error("failed to get system metrics", slog.String("error", err.Error()))
		return Metrics{}, err
	}
	n, err := net.IOCounters(false)
	if err != nil {
		slog.Error("failed to get system metrics", slog.String("error", err.Error()))
		return Metrics{}, err
	}
	c, err := cpu.Percent(time.Second, false)
	if err != nil {
		slog.Error("failed to get system metrics", slog.String("error", err.Error()))
		return Metrics{}, err
	}

	return Metrics{v.Used, n, c, time.Now().Unix()}, err
}

func (h *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	slog.Info("WebSocket connection established")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("failed to upgrade connection", slog.String("error", err.Error()))
		return
	}
	defer conn.Close()

	for {
		metrics, err := getSystemMetrics()
		if err != nil {
			slog.Error("failed to get system metrics", slog.String("error", err.Error()))
			break
		}

		err = conn.WriteJSON(metrics)
		if err != nil {
			slog.Error("failed to write JSON", slog.String("error", err.Error()))
			break
		}
		time.Sleep(1 * time.Second)
	}
}
