package authhandler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/AmadoMuerte/FlickSynergy/internal/db/repository"
	"github.com/AmadoMuerte/FlickSynergy/internal/jwt"
	"github.com/AmadoMuerte/FlickSynergy/internal/lib/response"
	"github.com/AmadoMuerte/FlickSynergy/internal/lib/validator"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
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
	userRepository := repository.NewUserRepository(h.db)
	user, err := userRepository.GetByEmail(req.Email)
	if err != nil {
		slog.Error("user not found", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusInternalServerError, "User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		slog.Error("invalid password", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusBadRequest, "invalid password")
		return
	}

	accessToken, err := jwt.GenerateToken(&jwt.UserInfo{ID: user.ID}, h.cfg.JWT.AcessDuration, h.cfg.JWT.Key, "access")
	if err != nil {
		slog.Error("failed to generate access token", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusInternalServerError, "failed to generate access token")
		return
	}
	refreshToken, err := jwt.GenerateToken(&jwt.UserInfo{ID: user.ID}, h.cfg.JWT.RefreshDuration, h.cfg.JWT.Key, "refresh")
	if err != nil {
		slog.Error("failed to generate refresh token", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusInternalServerError, "failed to generate refresh token")
		return
	}

	response.RespondWithJSON(w, r, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
