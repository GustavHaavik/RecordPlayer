package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"recordplayer/internal/domain"
)

type InMemory struct {
	albums []domain.Album
}

func NewInMemory(filepath string) *InMemory {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	albums := []domain.Album{}
	json.Unmarshal(file, &albums)

	return &InMemory{albums}
}

func (i *InMemory) GetAlbums() ([]domain.Album, error) {
	return i.albums, nil
}

func (i *InMemory) GetAlbumByID(id string) (domain.Album, error) {
	for _, a := range i.albums {
		if a.ID == id {
			return a, nil
		}
	}

	return domain.Album{}, errors.New("album not found")
}

func (i *InMemory) PostAlbum(album domain.Album) error {
	i.albums = append(i.albums, album)
	return nil
}
