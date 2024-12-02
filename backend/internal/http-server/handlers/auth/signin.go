package authhandler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/AmadoMuerte/sysStats/internal/db/repository"
	"github.com/AmadoMuerte/sysStats/internal/jwt"
	"github.com/AmadoMuerte/sysStats/internal/lib/response"
	"github.com/AmadoMuerte/sysStats/internal/lib/validator"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Sign In
// @Description This endpoint allows user to sign in using their email and passwd.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Credentials for signing in"
// @Success 200 {object} tokenResponse
// @Failure 401 {object} response.errorResponse
// @Failure 400 {object} response.errorResponse
// @Router /login/sign-in [post]
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
		slog.Error("Email or password is incorrect", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		slog.Error("Email or password is incorrect", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	accessToken, err := jwt.GenerateToken(&jwt.UserInfo{ID: user.ID}, h.cfg.JWT.AcessDuration, h.cfg.JWT.Key, "access")
	if err != nil {
		slog.Error("failed to generate access token", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusUnauthorized, "Email or password is incorrect")
		return
	}
	refreshToken, err := jwt.GenerateToken(&jwt.UserInfo{ID: user.ID}, h.cfg.JWT.RefreshDuration, h.cfg.JWT.Key, "refresh")
	if err != nil {
		slog.Error("failed to generate refresh token", slog.String("error", err.Error()))
		response.RespondWithError(w, r, http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	response.RespondWithJSON(w, r, http.StatusOK, &tokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
