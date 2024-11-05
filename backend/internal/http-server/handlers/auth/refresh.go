package authhandler

import (
	"net/http"
	"strings"

	"github.com/AmadoMuerte/FlickSynergy/internal/jwt"
	"github.com/AmadoMuerte/FlickSynergy/internal/lib/response"
	"github.com/go-chi/render"
)

// @Summary Refresh Access Token
// @Description This endpoint allows users to refresh their access token using a valid refresh token.
//
//	The client must provide the refresh token in the Authorization header as a Bearer token.
//
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {refresh_token}"
// @Success 200 {object} refreshResponse
// @Failure 401 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /login/refresh [post]
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid refresh token format")
		return
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid refresh token format")
		return
	}

	refreshToken := tokenParts[1]
	_, err := jwt.VerifyToken(refreshToken, h.cfg.JWT.Key, "refresh")
	if err != nil {
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	userInfo, err := jwt.ExtractUserInfo(refreshToken, []byte(h.cfg.JWT.Key))
	if err != nil {
		response.RespondWithError(w, r, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	newAccessToken, err := jwt.GenerateToken(userInfo, h.cfg.JWT.AcessDuration, h.cfg.JWT.Key, "access")
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "could not generate access token")
		return
	}

	newRefreshToken, err := jwt.GenerateToken(userInfo, h.cfg.JWT.RefreshDuration, h.cfg.JWT.Key, "refresh")
	if err != nil {
		response.RespondWithError(w, r, http.StatusInternalServerError, "could not generate refresh token")
		return
	}

	render.JSON(w, r, refreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}

type refreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
