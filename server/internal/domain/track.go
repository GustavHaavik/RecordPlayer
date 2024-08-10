package domain

type Track struct {
	Title    string `json:"title"`
	Duration int    `json:"length"`
	AlbumID  string `json:"album_id"`
}
