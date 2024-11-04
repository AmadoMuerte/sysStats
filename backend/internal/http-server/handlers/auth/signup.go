package authhandler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/AmadoMuerte/FlickSynergy/internal/db/models"
	"github.com/AmadoMuerte/FlickSynergy/internal/db/repository"
	"github.com/AmadoMuerte/FlickSynergy/internal/http-server/handlers"
	"github.com/AmadoMuerte/FlickSynergy/internal/jwt"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, r, http.StatusBadRequest, "failed to decode data")
		return
	}

	if !validateEmail(req.Email) {
		h.respondWithError(w, r, http.StatusBadRequest, "invalid email")
		return
	}

	if err := validatePassword(req.Password); err != nil {
		h.respondWithError(w, r, http.StatusBadRequest, "invalid password")
		return
	}

	user := &models.User{Email: req.Email, Password: req.Password}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", slog.String("error", err.Error()))
		h.respondWithError(w, r, http.StatusInternalServerError, "failed to create user")
		return
	}
	user.Password = string(hashedPassword)

	userRepository := repository.NewUserRepository(h.db)
	userID, err := userRepository.Create(user)
	if err != nil {
		h.respondWithError(w, r, http.StatusInternalServerError, "failed to create user")
		return
	}

	refreshToken, err := jwt.GenerateToken(&jwt.UserInfo{ID: userID}, h.cfg.JWT.RefreshDuration, h.cfg.JWT.Key, "refresh")
	if err != nil {
		h.respondWithError(w, r, http.StatusInternalServerError, "failed to create refresh token")
		return
	}

	accessToken, err := jwt.GenerateToken(&jwt.UserInfo{ID: userID}, h.cfg.JWT.AcessDuration, h.cfg.JWT.Key, "access")
	if err != nil {
		h.respondWithError(w, r, http.StatusInternalServerError, "failed to create access token")
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, struct {
		Status       int    `json:"status"`
		Message      string `json:"message"`
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}{
		Status:       http.StatusCreated,
		Message:      "user created",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) respondWithError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	render.JSON(w, r, handlers.ErrorResponse{
		Status:  status,
		Message: message,
	})
}
