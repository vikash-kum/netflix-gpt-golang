package router

import (
	"net/http"
	controllers "netflix-gpt-backend/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Set the allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Message": "Hello World"})
	})
	// Protected routes
	auth := router.Group("/users").Use(AuthMiddleware())
	{
		auth.GET("/", controllers.GetAllUsers)
		auth.GET("/:id", controllers.GetUser)
		auth.PUT("/:id", controllers.UpdateUser)
		auth.DELETE("/:id", controllers.DeleteUser)
	}

	imdb := router.Group("/imdb")
	{
		imdb.GET("/nowPlaying", controllers.GetNowPlaying)
		imdb.GET("/movie_trailer/:movieId", controllers.GetMovieTrailer)
		imdb.GET("/popular_movies", controllers.GetPopularMovies)
		imdb.GET("/top_rated_movies", controllers.GetTopRatedMovies)
		imdb.GET("/up_coming_movies", controllers.GetUpcomingMovies)
	}

	return router
}
