package main

import (
	"recordplayer/internal/adapters/api"
	"recordplayer/internal/adapters/db"
	"recordplayer/internal/adapters/services"
)

func main() {
	albumRepo := db.NewInMemory("../server/albums.json")

	service := services.NewAlbumService(albumRepo)

	api := api.NewApiServer(service)
	api.Start(":8080")
}
