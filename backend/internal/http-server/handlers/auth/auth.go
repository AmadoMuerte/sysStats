package authhandler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/db"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func validatePassword(pass string) error {
	var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()-_+=]{8,20}$`)

	if !passwordRegex.MatchString(pass) {
		return errors.New("invalid username or password")
	}

	return nil
}

func validateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return emailRegex.MatchString(email)
}
