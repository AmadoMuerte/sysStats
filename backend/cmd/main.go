package main

import (
	"fmt"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/db"
	"github.com/AmadoMuerte/FlickSynergy/internal/http-server/server"
)

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
