package authhandler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/AmadoMuerte/FlickSynergy/internal/db/models"
	"github.com/AmadoMuerte/FlickSynergy/internal/db/repository"
	"github.com/AmadoMuerte/FlickSynergy/internal/lib/response"
	"github.com/AmadoMuerte/FlickSynergy/internal/lib/validator"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode data", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusBadRequest, "invalid data")
		return
	}

	if !validator.ValidateEmail(req.Email) {
		slog.Error("invalid email", slog.String("error", "invalid email"))
		response.RespondWithError(w, r, http.StatusBadRequest, "invalid email")
		return
	}

	if err := validator.ValidatePassword(req.Password); err != nil {
		slog.Error("invalid password", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusBadRequest, "invalid password")
		return
	}

	user := &models.User{Email: req.Email, Password: req.Password}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusInternalServerError, "failed to hash password")
		return
	}
	user.Password = string(hashedPassword)

	userRepository := repository.NewUserRepository(h.db)
	_, err = userRepository.Create(user)
	if err != nil {
		slog.Error("failed to create user", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusInternalServerError, "Email already exists")
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  http.StatusCreated,
		Message: "user created, please login",
	})
}
