package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://go.dev/doc/tutorial/web-service-gin
// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "The Dave Brubeck Quartet", Artist: "Dave Brubeck", Price: 26.99},
	{ID: "2", Title: "1994 Oxygene [Trance Remix]", Artist: "Jean Michel Jarre", Price: 17.99},
	{ID: "3", Title: "Riding With The King(2000)", Artist: "Eric Clapton & B.B. King", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*
try it with
$ curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

+ GIN: Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
*/
