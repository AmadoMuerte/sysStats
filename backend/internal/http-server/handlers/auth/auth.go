package authhandler

import (
	"net/http"

	"github.com/AmadoMuerte/sysStats/internal/config"
	"github.com/AmadoMuerte/sysStats/internal/db"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthHandler struct {
	cfg *config.Config
	db  *db.Storage
}

type IAuth interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

func New(cfg *config.Config, db *db.Storage) IAuth {
	return &AuthHandler{cfg, db}
}
