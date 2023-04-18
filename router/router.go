package router

import (
	"api/docs"
	"api/handler/api"
	"api/middleware"
	"api/sse"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
)

func InitRouter(app *gin.Engine, appHandler api.AppHandler, db *gorm.DB, stream *sse.Event) *gin.Engine {
	middlewareFunc := middleware.NewJwtMiddleware(db)

	app.GET("/health_check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"service_name": "goldenowl_gin_api",
			"status":       "ok",
			"data":         "📺 API Up and Running",
		})
	})

	usersV1 := app.Group("api/v1")
	{
		usersV1.POST("user", appHandler.User.Create)
		usersV1.GET("user", middlewareFunc.DeserializeUser(), appHandler.User.GetByUsername)
		usersV1.POST("user/login", appHandler.User.Login)
		usersV1.POST("user/logout", middlewareFunc.DeserializeUser(), appHandler.User.Logout)
	}

	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.StaticFS("/images", http.Dir("public"))

	return app
}
