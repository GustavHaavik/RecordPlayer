package ports

import "recordplayer/internal/domain"

type AlbumRepository interface {
	GetAlbums() ([]domain.Album, error)
	GetAlbumByID(id string) (domain.Album, error)
	PostAlbum(album domain.Album) error
}
