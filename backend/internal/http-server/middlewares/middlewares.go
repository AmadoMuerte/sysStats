package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/jwt"
	"github.com/go-chi/render"
)

type AuthMiddleware struct {
	Cfg *config.Config
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (m *AuthMiddleware) New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			respondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		token := tokenParts[1]

		_, err := jwt.VerifyToken(token, m.Cfg.JWT.Key)
		if err != nil {
			respondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		userInfo, err := jwt.ExtractUserInfo(token, []byte(m.Cfg.JWT.Key))
		if err != nil {
			respondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		ctx := context.WithValue(r.Context(), userInfo, userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func respondWithError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	render.JSON(w, r, response{
		Status:  status,
		Message: message,
	})
}
