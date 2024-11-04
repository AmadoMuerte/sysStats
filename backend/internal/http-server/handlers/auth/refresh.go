package authhandler

import (
	"net/http"
	"strings"

	"github.com/AmadoMuerte/FlickSynergy/internal/http-server/handlers"
	"github.com/AmadoMuerte/FlickSynergy/internal/jwt"
	"github.com/go-chi/render"
)

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, handlers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "refresh token not provided",
		})
		return
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, handlers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "invalid refresh token format",
		})
		return
	}

	refreshToken := tokenParts[1]
	_, err := jwt.VerifyToken(refreshToken, h.cfg.JWT.Key, "refresh")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, handlers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "invalid refresh token",
		})
		return
	}

	userInfo, err := jwt.ExtractUserInfo(refreshToken, []byte(h.cfg.JWT.Key))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, handlers.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "invalid refresh token",
		})
		return
	}

	// Генерируем новые токены
	newAccessToken, err := jwt.GenerateToken(userInfo, h.cfg.JWT.AcessDuration, h.cfg.JWT.Key, "access")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, handlers.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "could not generate access token",
		})
		return
	}

	newRefreshToken, err := jwt.GenerateToken(userInfo, h.cfg.JWT.RefreshDuration, h.cfg.JWT.Key, "refresh")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, handlers.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "could not generate refresh token",
		})
		return
	}

	render.JSON(w, r, struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}
