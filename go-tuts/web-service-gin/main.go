package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct { // declaring types for each prop of album
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// getAlbums responds with the list fo all albums as JSON
func getAlbums(c *gin.Context) { // context carries the request details
	c.IndentedJSON(http.StatusOK, albums)
	// serialize the struct into JSOn and add it to the response
	// first arg: status code to be returned
	// second arg: data to be serialized
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()                 // establish a gin router using default
	router.GET("/albums", getAlbums)        // associate get function with the /albums path. NOTE: no need for () just the name will do
	router.GET("/albums/:id", getAlbumByID) // if colon, represents path parameter
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080") // Attach the router to a server
}

// postAlbums adds an album from JSOn received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSOn to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id") // get id from param

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
