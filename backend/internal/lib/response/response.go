package response

import (
	"net/http"

	"github.com/go-chi/render"
)

// errorResponse represents the structure of an error response returned to the client.
// @Description This structure is used to send error information in JSON format.
// @Property Status int `json:"status"` "HTTP status code indicating the error type"
// @Property Message string `json:"message"` "Detailed error message explaining the issue"
type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// RespondWithError sends a JSON response with the specified error status and message.
// @Summary Respond with Error
// @Description This function writes an HTTP response with the specified status code and error message in JSON format.
// @Param status query int true "HTTP status code"
// @Param message query string true "Error message"
// @Success 200 {object} errorResponse "Successful error response"
// @Failure 400 {object} errorResponse "Bad request error"
// @Failure 500 {object} errorResponse "Internal server error"
func RespondWithError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	render.JSON(w, r, errorResponse{
		Status:  status,
		Message: message,
	})
}

// RespondWithJSON sends a JSON response with the specified status code and data.
// @Summary Respond with JSON
// @Description This function writes an HTTP response with the specified status code and data in JSON format.
// @Param status query int true "HTTP status code"
// @Param v body interface{} true "Response data"
// @Success 200 {object} interface{} "Successful JSON response"
// @Failure 400 {object} errorResponse "Bad request error"
// @Failure 500 {object} errorResponse "Internal server error"
func RespondWithJSON(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	w.WriteHeader(status)
	render.JSON(w, r, v)
}
