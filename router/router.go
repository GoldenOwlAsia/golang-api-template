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
	middlewareFunc := middleware.NewJwtMiddleware(db)

	app.GET("/health_check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"service_name": "Golden Owl Gin API",
			"status":       "ok",
			"data":         "📺 API up and running",
		})
	})

	users := app.Group("api/v1/user")
	users.POST("/login", appHandler.User.Login)
	users.POST("/generateToken", middlewareFunc.DeserializeUser(), appHandler.User.GenerateTokenHandler)
	users.POST("/refreshAccessToken", middlewareFunc.DeserializeUser(), appHandler.User.RefreshAccessTokenHandler)

	articles := app.Group("api/v1/articles")
	articles.GET("/", middlewareFunc.DeserializeUser(), appHandler.Article.All)
	articles.GET("/:id", middlewareFunc.DeserializeUser(), appHandler.Article.Get)
	articles.POST("/", middlewareFunc.DeserializeUser(), appHandler.Article.Create)
	articles.PUT("/:id", middlewareFunc.DeserializeUser(), appHandler.Article.Update)
	articles.DELETE("/:id", middlewareFunc.DeserializeUser(), appHandler.Article.Delete)

	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.StaticFS("/images", http.Dir("public"))

	return app
}
