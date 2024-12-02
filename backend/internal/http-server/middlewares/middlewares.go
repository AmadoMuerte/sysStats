package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/AmadoMuerte/sysStats/internal/config"
	"github.com/AmadoMuerte/sysStats/internal/jwt"
	"github.com/AmadoMuerte/sysStats/internal/lib/response"
)

type AuthMiddleware struct {
	Cfg *config.Config
}

func (m *AuthMiddleware) New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.RespondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			response.RespondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		token := tokenParts[1]

		_, err := jwt.VerifyToken(token, m.Cfg.JWT.Key, "access")
		if err != nil {
			response.RespondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		userInfo, err := jwt.ExtractUserInfo(token, []byte(m.Cfg.JWT.Key))
		if err != nil {
			response.RespondWithError(w, r, http.StatusUnauthorized, "token not valid")
			return
		}

		ctx := context.WithValue(r.Context(), userInfo, userInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
