package main

import (
	"fmt"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/db"
	"github.com/AmadoMuerte/FlickSynergy/internal/db/models"
	"github.com/AmadoMuerte/FlickSynergy/internal/db/repository"
)

func main() {
	cfg, err := config.NewConfig(nil)
	if err != nil {
		panic(err)
	}

	storage := db.New(cfg)
	fmt.Print("Connected to DB\n")

	userRepo := repository.NewUserRepository(storage)
	userRepo.Create(&models.User{
		Username: "test",
		Email:    "test",
	})
}
