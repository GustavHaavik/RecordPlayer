package api

import (
	"recordplayer/internal/domain"
	"recordplayer/internal/ports"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	albumService ports.AlbumService
}

func NewApiServer(albumService ports.AlbumService) *ApiServer {
	return &ApiServer{albumService}
}

func (s *ApiServer) Start(listenAddr string) error {
	router := gin.Default()

	router.GET("/albums", s.getAlbums)
	router.GET("/albums/:id", s.getAlbumByID)
	router.POST("/albums", s.postAlbum)

	if err := router.Run(listenAddr); err != nil {
		return err
	}

	return nil
}

func (s *ApiServer) getAlbums(c *gin.Context) {
	albums, err := s.albumService.GetAlbums()
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(200, albums)
}

func (s *ApiServer) getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := s.albumService.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(200, album)
}

func (s *ApiServer) postAlbum(c *gin.Context) {
	var newAlbum domain.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request"})
		return
	}

	if err := s.albumService.PostAlbum(newAlbum); err != nil {
		c.IndentedJSON(500, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(201, newAlbum)
}
