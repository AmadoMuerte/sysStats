package authhandler

import (
	"net/http"
	"strings"

	"github.com/AmadoMuerte/sysStats/internal/jwt"
	"github.com/AmadoMuerte/sysStats/internal/lib/response"
	"github.com/go-chi/render"
)

// @Summary Refresh Access token
// @Description This endpoint allows users to refresh their access token using a valid refresh token.
//
//	The client must provide the refresh token in the Authorization header as a Bearer token.
//
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {refresh_token}"
// @Success 200 {object} tokenResponse
// @Failure 401 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /login/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid token")
		return
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid token")
		return
	}

	refreshToken := tokenParts[1]
	_, err := jwt.VerifyToken(refreshToken, h.cfg.JWT.Key, "refresh")
	if err != nil {
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid token")
		return
	}

	userInfo, err := jwt.ExtractUserInfo(refreshToken, []byte(h.cfg.JWT.Key))
	if err != nil {
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid token")
		return
	}

	newAccessToken, err := jwt.GenerateToken(userInfo, h.cfg.JWT.AcessDuration, h.cfg.JWT.Key, "access")
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	newRefreshToken, err := jwt.GenerateToken(userInfo, h.cfg.JWT.RefreshDuration, h.cfg.JWT.Key, "refresh")
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	render.JSON(w, r, tokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}
