package metric

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(URL, h.Heartbeat)
}

// Heartbeat
// @Summary Heartbeat metric
// @Tags metrics
// @Succes 204
// @Failure 400
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	log.Print("HearthBeat")
	w.WriteHeader(204)
}
