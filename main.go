package main

import (
	"api/api"
	"api/infras"
	"api/middleware"
	"api/migrations"
	"api/router"
	sentrygin "github.com/getsentry/sentry-go/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func init() {
	api.Api = api.NewServer()
	api.Api.InitEnv().InitDb().InitGin().InitSentry().InitSSE()
}

// @title           GoldenOwl Gin API
// @version         1.0.0
// @description     This API is for GoldenOwl API application
// @contact.name	joe.nghiem
// @contact.email	joe.nghiem.goldenowl@gmail.com
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
// @securityDefinitions.basic  BasicAuth
func main() {
	migrations.Migrate()
	migrations.Seed()
	api.Api.Gin.Use(middleware.Cors())
	api.Api.Gin.Use(middleware.HandleResponse)
	api.Api.Gin.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	appHandler := infras.DI(api.Api.DB, api.Api.Event)
	api.Api.Gin = router.InitRouter(api.Api.Gin, appHandler, api.Api.DB, api.Api.Event)
	api.Api.SpinUp()
}
