package main

import (
	"fmt"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/db"
	"github.com/AmadoMuerte/FlickSynergy/internal/http-server/server"
)

// @title FlickSynergy API
// @version 1.0
// @description API для FlickSynergy, платформы для сбора предпочтений пользователей при выборе фильма для совместного просмотра.
// @contact.name GitHub
// @contact.url https://github.com/AmadoMuerte
// @basePath /api/v1
// @schemes http https
func main() {
	cfg, err := config.NewConfig(nil)
	if err != nil {
		panic(err)
	}

	storage := db.New(cfg)
	fmt.Print("Connected to DB\n")

	server := server.New(cfg, storage)
	server.Start()
}
