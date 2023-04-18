package main

import (
	"api/migrations"
	"api/sse"
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	gormlib "gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"

	"api/configs"
	"api/handler/api"
	"api/middleware"
	"api/pkgs/gorm"
	"api/router"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

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
	var err error
	// load .env config
	_, err = configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs: ", err)
	}
	if err = os.Setenv("TZ", configs.ConfApp.AppTimezone); err != nil {
		log.Fatal("cannot set TZ config: ", err)
	}

	// init sentry
	if err = sentry.Init(sentry.ClientOptions{
		Dsn:              configs.ConfApp.SentryDNS,
		EnableTracing:    true,
		Environment:      configs.ConfApp.AppMode,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}

	var (
		app    = gin.Default()
		stream = sse.NewServer()
		db     = gorm.DbInstance.Instance()
	)

	gin.SetMode(configs.ConfApp.AppMode)
	migrations.Migrate()

	app.Use(middleware.Cors())
	app.Use(middleware.HandleResponse)
	app.Use(sentrygin.New(sentrygin.Options{Repanic: true}))

	appHandler := DI(db, stream)

	app = router.InitRouter(app, appHandler, db, stream)

	// -- Graceful restart or stop server --
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.ConfApp.Port), // + configs.ConfApp.Port
		Handler: app,
	}

	fmt.Printf("[GoldenOwl API - %s] Start to listening the incoming requests on http address: %s 🚀🚀🚀", gin.Mode(), server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}

func DI(db *gormlib.DB, event *sse.Event) api.AppHandler {
	return api.AppHandler{
		User: InitUserAPI(db),
	}
}
