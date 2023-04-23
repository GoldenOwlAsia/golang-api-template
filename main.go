package main

import (
	"fmt"
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/infras"
	"github.com/GoldenOwlAsia/golang-api-template/middleware"
	"github.com/GoldenOwlAsia/golang-api-template/migrations"
	"github.com/GoldenOwlAsia/golang-api-template/router"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	DB  *gorm.DB
	Gin *gin.Engine
)

// @title           GoldenOwl Gin API
// @version         1.0.0
// @description     This API is for GoldenOwl API application
// @contact.name	goldenowl.asia
// @contact.email	hello@goldenowl.asia
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
// @securityDefinitions.basic  BasicAuth
func main() {
	InitEnv()
	InitDb()
	InitGin()
	InitSentry()
	migrations.Migrate(DB)
	migrations.Seed(DB)
	Gin.Use(middleware.Cors())
	Gin.Use(middleware.HandleResponse)
	Gin.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	appHandler := infras.DI(DB)
	Gin = router.InitRouter(Gin, appHandler, DB)

	SpinUp(" 🚀 ")
}

func InitEnv() {
	var err error
	_, err = configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs: ", err)
	}
	if err = os.Setenv("TZ", configs.ConfApp.AppTimezone); err != nil {
		log.Fatal("cannot set TZ config: ", err)
	}
}

func InitDb() {
	dsn := getDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if configs.ConfApp.AppMode == gin.DebugMode {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(),
	})

	if err != nil {
		log.Println("cannot connect to database: ", err)
	}

	DB = db

}

func getDSN() (dsn string) {
	username := configs.ConfApp.DatabaseUserName
	password := configs.ConfApp.DatabasePassword
	host := configs.ConfApp.DatabaseHost
	port := configs.ConfApp.DatabasePort
	databaseName := configs.ConfApp.DatabaseName
	timeZone := configs.ConfApp.DatabaseTimezone
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, username, password, databaseName, port, timeZone)

	return
}

func InitGin() {
	g := gin.Default()
	Gin = g
	gin.SetMode(configs.ConfApp.AppMode)
}

func InitSentry() {
	var err error
	if err = sentry.Init(sentry.ClientOptions{
		Dsn:              configs.ConfApp.SentryDNS,
		EnableTracing:    true,
		Environment:      configs.ConfApp.AppMode,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}
}

func SpinUp(msg string) (server *http.Server) {
	// -- Graceful restart or stop server --
	server = &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.ConfApp.Port), // + configs.ConfApp.Port
		Handler: Gin,
	}
	fmt.Printf("[GoldenOwl API - %s] Start to listening the incoming requests on http address: %s", gin.Mode(), server.Addr)
	fmt.Println(msg)

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

	return
}
