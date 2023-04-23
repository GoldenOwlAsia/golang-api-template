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

	usersV1 := app.Group("api/v1/user")
	{
		usersV1.POST("/", appHandler.User.Create)
		usersV1.GET("/", middlewareFunc.DeserializeUser(), appHandler.User.GetByUsername)
		usersV1.POST("/login", appHandler.User.Login)
		usersV1.POST("/logout", middlewareFunc.DeserializeUser(), appHandler.User.Logout)
		usersV1.POST("/generateToken", middlewareFunc.DeserializeUser(), appHandler.User.GenerateTokenHandler)
		usersV1.POST("/refreshAccessToken", middlewareFunc.DeserializeUser(), appHandler.User.RefreshAccessTokenHandler)
	}

	articlesAPI := app.Group("api/v1/articles")
	{
		articlesAPI.GET("/", middlewareFunc.DeserializeUser(), appHandler.Article.All)
		articlesAPI.GET("/:id", middlewareFunc.DeserializeUser(), appHandler.Article.Get)
		articlesAPI.POST("/", middlewareFunc.DeserializeUser(), appHandler.Article.Create)
		articlesAPI.PUT("/:id", middlewareFunc.DeserializeUser(), appHandler.Article.Update)
		articlesAPI.DELETE("/:id", middlewareFunc.DeserializeUser(), appHandler.Article.Delete)
	}

	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.StaticFS("/images", http.Dir("public"))

	return app
}
