package services

import (
	"recordplayer/internal/domain"
	"recordplayer/internal/ports"
)

type AlbumService struct {
	albumRepo ports.AlbumRepository
}

func NewAlbumService(albumRepo ports.AlbumRepository) *AlbumService {
	return &AlbumService{albumRepo}
}

func (s *AlbumService) GetAlbums() ([]domain.Album, error) {
	return s.albumRepo.GetAlbums()
}

func (s *AlbumService) GetAlbumByID(id string) (domain.Album, error) {
	return s.albumRepo.GetAlbumByID(id)
}

func (s *AlbumService) PostAlbum(album domain.Album) error {
	return s.albumRepo.PostAlbum(album)
}
