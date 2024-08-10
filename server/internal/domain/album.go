package domain

type Album struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	CoverArt string  `json:"cover_art"`
	Year     string  `json:"year"`
	Price    float32 `json:"price"`
	// Tracks   []Track `json:"tracks"`
}
