package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	render.JSON(w, r, errorResponse{
		Status:  status,
		Message: message,
	})
}
