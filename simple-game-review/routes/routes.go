package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gamereview/controller"
	"gamereview/middleware"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	r.GET("/games", controller.GetGames)
	r.GET("/:id", controller.GetGameID)
	r.GET("/games/:id/reviews", controller.GetReviewByGameId)

	gamesMiddlewareRoute := r.Group("/games")
	gamesMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	gamesMiddlewareRoute.POST("/", controller.CreateGame)
	gamesMiddlewareRoute.PATCH("/:id", controller.UpdateGame)
	gamesMiddlewareRoute.DELETE("/:id", controller.DeleteGame)

	r.GET("/developer", controller.GetDevs)
	r.GET("/developer/:id", controller.GetDevsID)
	r.GET("/developer/:id/games", controller.GetGamesByDevId)

	developerMiddlewareRoute := r.Group("/developer")
	developerMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	developerMiddlewareRoute.POST("/", controller.CreateDevs)
	developerMiddlewareRoute.PATCH("/:id", controller.UpdateDev)
	developerMiddlewareRoute.DELETE("/:id", controller.DeleteDev)

	r.GET("/publisher", controller.GetPub)
	r.GET("/publisher/:id", controller.GetPubID)
	r.GET("/publisher/:id/games", controller.GetGamesByPubId)

	publisherMiddlewareRoute := r.Group("/publisher")
	publisherMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	publisherMiddlewareRoute.POST("/", controller.CreatePubs)
	publisherMiddlewareRoute.PATCH("/:id", controller.UpdatePub)
	publisherMiddlewareRoute.DELETE("/:id", controller.DeletePub)

	r.GET("/category", controller.GetCategory)
	r.GET("/category/:id", controller.GetCategoryID)
	r.GET("/category/:id/games", controller.GetGamesByCategoryId)

	categoryMiddlewareRoute := r.Group("/category")
	categoryMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	categoryMiddlewareRoute.POST("/", controller.CreateCategory)
	categoryMiddlewareRoute.PATCH("/:id", controller.UpdateCategory)
	categoryMiddlewareRoute.DELETE("/:id", controller.DeleteCategory)

	r.GET("/review", controller.GetReview)
	r.GET("/review/:id", controller.GetReviewID)
	r.GET("/review/:id/ratings", controller.GetRatingByReviewId)

	reviewMiddlewareRoute := r.Group("/review")
	reviewMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	reviewMiddlewareRoute.POST("/", controller.CreateReview)
	reviewMiddlewareRoute.PATCH("/:id", controller.UpdateReview)
	reviewMiddlewareRoute.DELETE("/:id", controller.DeleteReview)

	r.GET("/rating", controller.GetRate)
	r.GET("/rating/:id", controller.GetRateID)

	ratingMiddlewareRoute := r.Group("/rating")
	ratingMiddlewareRoute.Use(middleware.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("/", controller.CreateRate)
	ratingMiddlewareRoute.PATCH("/:id", controller.UpdateRate)
	ratingMiddlewareRoute.DELETE("/:id", controller.DeleteRate)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
