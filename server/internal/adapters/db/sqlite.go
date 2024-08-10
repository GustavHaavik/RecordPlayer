package db

import (
	"database/sql"
	"log"
	"recordplayer/internal/domain"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLite() *SQLiteRepository {
	db, err := sql.Open("sqlite3", "database/items.db")
	if err != nil {
		log.Fatal("Couldn't initialize sqlite db")
	}
	_, err = db.Exec("create table if not exists items (id string, name text, description text)")
	if err != nil {
		log.Fatal("Couldn't create initial items table")
	}
	return &SQLiteRepository{db}
}

func (s *SQLiteRepository) GetAlbums() ([]domain.Album, error) {
	rows, err := s.db.Query("select * from albums")
	if err != nil {
		return nil, err
	}

	albums := []domain.Album{}
	for rows.Next() {
		var album domain.Album
		rows.Scan(&album.ID, &album.Title, &album.Artist, &album.CoverArt, &album.Year, &album.Price)
		albums = append(albums, album)
	}

	return albums, nil
}

func (s *SQLiteRepository) GetAlbumByID(id string) (domain.Album, error) {
	row := s.db.QueryRow("select * from albums where id = ?", id)
	var album domain.Album
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.CoverArt, &album.Year, &album.Price)
	if err != nil {
		return domain.Album{}, err
	}

	return album, nil
}

func (s *SQLiteRepository) PostAlbum(album domain.Album) error {
	_, err := s.db.Exec("insert into albums (id, title, artist, cover_art, year, price) values (?, ?, ?, ?, ?, ?)", album.ID, album.Title, album.Artist, album.CoverArt, album.Year, album.Price)
	if err != nil {
		return err
	}

	return nil
}
