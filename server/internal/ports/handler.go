package ports

import "recordplayer/internal/domain"

type AlbumHandler interface {
	GetAlbums() ([]domain.Album, error)
	GetAlbumByID(id string) (domain.Album, error)
	PostAlbum(album domain.Album) error
}
