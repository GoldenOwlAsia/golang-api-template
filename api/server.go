package api

import (
	"api/configs"
	"api/sse"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	DB    *gorm.DB
	Gin   *gin.Engine
	Event *sse.Event
}

var (
	Api *Server
)

func NewServer() *Server {
	return &Server{}
}

func (s *Server) InitEnv() *Server {
	var err error
	_, err = configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs: ", err)
	}
	if err = os.Setenv("TZ", configs.ConfApp.AppTimezone); err != nil {
		log.Fatal("cannot set TZ config: ", err)
	}
	return s
}

func (s *Server) InitDb() *Server {
	dsn := s.getDSN()
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

	s.DB = db

	return s
}

func (s *Server) getDSN() (dsn string) {
	username := configs.ConfApp.DatabaseUserName
	password := configs.ConfApp.DatabasePassword
	host := configs.ConfApp.DatabaseHost
	port := configs.ConfApp.DatabasePort
	databaseName := configs.ConfApp.DatabaseName
	timeZone := configs.ConfApp.DatabaseTimezone
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, username, password, databaseName, port, timeZone)

	return
}

func (s *Server) InitGin() *Server {
	g := gin.Default()
	s.Gin = g
	gin.SetMode(configs.ConfApp.AppMode)
	return s
}

func (s *Server) InitSSE() *Server {
	s.Event = sse.NewServer()

	return s
}

func (s *Server) InitSentry() *Server {
	var err error
	if err = sentry.Init(sentry.ClientOptions{
		Dsn:              configs.ConfApp.SentryDNS,
		EnableTracing:    true,
		Environment:      configs.ConfApp.AppMode,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
	}
	return s
}

func (s *Server) SpinUp() *Server {
	// -- Graceful restart or stop server --
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.ConfApp.Port), // + configs.ConfApp.Port
		Handler: Api.Gin,
	}

	fmt.Printf("[GoldenOwl API - %s] Start to listening the incoming requests on http address: %s 🚀", gin.Mode(), server.Addr)

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

	return s
}
