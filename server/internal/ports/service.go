package ports

import "recordplayer/internal/domain"

type AlbumService interface {
	GetAlbums() ([]domain.Album, error)
	GetAlbumByID(id string) (domain.Album, error)
	PostAlbum(album domain.Album) error
}
