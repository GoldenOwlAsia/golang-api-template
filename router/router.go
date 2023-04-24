package router

import (
	"github.com/GoldenOwlAsia/golang-api-template/docs"
	"github.com/GoldenOwlAsia/golang-api-template/infras"
	"github.com/GoldenOwlAsia/golang-api-template/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
)

func InitRouter(app *gin.Engine, appHandler infras.AppHandler, db *gorm.DB) *gin.Engine {
	jwtAuthMiddleware := middleware.NewJwtAuth(db)

	app.GET("/health_check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"service_name": "Golden Owl Gin API",
			"status":       "ok",
			"data":         "📺 API up and running",
		})
	})

	users := app.Group("api/v1/user")
	users.POST("/login", appHandler.User.Login)
	users.POST("/refreshAccessToken", appHandler.User.RefreshAccessToken)

	articles := app.Group("api/v1/articles", jwtAuthMiddleware.Authenticate())
	articles.GET("/", appHandler.Article.All)
	articles.GET("/:id", appHandler.Article.Get)
	articles.POST("/", appHandler.Article.Create)
	articles.PUT("/:id", appHandler.Article.Update)
	articles.DELETE("/:id", appHandler.Article.Delete)

	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.StaticFS("/images", http.Dir("public"))

	return app
}
