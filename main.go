package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumsById)
	router.POST("/albums", PostAlbums)


	router.Run("localhost:8080")
}

func PostAlbums(c *gin.Context) {
	var newAlbums album

	if err := c.BindJSON(&newAlbums); err != nil {
		return
	}
	albums = append(albums, newAlbums)
	c.IndentedJSON(http.StatusCreated, newAlbums)

}

func GetAlbumsById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}
