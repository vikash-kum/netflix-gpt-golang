package controllers

import (
	"encoding/json"
	"net/http"
	services "netflix-gpt-backend/service" // Update with actual path

	"github.com/gin-gonic/gin"
)

// Injecting service for API calls
var imdbService = services.NewImdbService("Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyNzkwN2ZkZDczOWM4OGVlYTVkOTMyM2U4MDJlYzIzNSIsInN1YiI6IjY1NjZkOGUyMTU2Y2M3MDE0ZTY2NjQzOSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.xqSzX87E-qH58F9bjogjfq-CNh4GBd23z8Vqv3ZCSFw") // Replace with your actual TMDB token

func GetNowPlaying(ctx *gin.Context) {
	url := "https://api.themoviedb.org/3/movie/now_playing?page=1"
	response, err := imdbService.NowPlayingTMDBApi(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetMovieTrailer(ctx *gin.Context) {
	movieId := ctx.Param("movieId")
	url := "https://api.themoviedb.org/3/movie/" + movieId + "/videos?language=en-US"
	response, err := imdbService.NowPlayingTMDBApi(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetPopularMovies(ctx *gin.Context) {
	url := "https://api.themoviedb.org/3/movie/popular?page=1"
	response, err := imdbService.NowPlayingTMDBApi(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetTopRatedMovies(ctx *gin.Context) {
	url := "https://api.themoviedb.org/3/movie/top_rated?page=1"
	response, err := imdbService.NowPlayingTMDBApi(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetUpcomingMovies(ctx *gin.Context) {
	url := "https://api.themoviedb.org/3/movie/upcoming?page=1"
	response, err := imdbService.NowPlayingTMDBApi(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
