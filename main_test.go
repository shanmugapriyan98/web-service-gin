package main

import (
	"example/web-service-gin/controllers"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAlbums(t *testing.T) {
	r := gin.Default()
	r.GET("/albums", controllers.GetAlbums)

	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// fmt.Println(http.StatusOK)
	// fmt.Println(w.Code)
	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `[
		{
			"id": "1",
			"title": "Blue Train",
			"artist": "John Coltrane",
			"price": 56.99
		},
		{
			"id": "2",
			"title": "Jeru",
			"artist": "Gerry Mulligan",
			"price": 17.99
		},
		{
			"id": "3",
			"title": "Sarah Vaughan and Clifford Brown",
			"artist": "Sarah Vaughan",
			"price": 39.99
		}
	]`

	response, _ := io.ReadAll(w.Body)
	// fmt.Print(expectedResponse)
	// fmt.Print(string(response))
	require.JSONEq(t, expectedResponse, string(response))
	assert.Equal(t, http.StatusOK, w.Code)

}
