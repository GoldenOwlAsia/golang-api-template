package gorm

import (
	"api/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	initializer func() *gorm.DB
	instance    *gorm.DB
	once        sync.Once
}

var (
	DbInstance *DBInstance
)

func (i *DBInstance) Instance() *gorm.DB {
	i.once.Do(func() {
		i.instance = i.initializer()
	})

	return i.instance
}

func init() {
	DbInstance = &DBInstance{initializer: dbInit}
}

func dbInit() *gorm.DB {
	username := configs.ConfApp.DatabaseUserName
	password := configs.ConfApp.DatabasePassword
	host := configs.ConfApp.DatabaseHost
	port := configs.ConfApp.DatabasePort
	databaseName := configs.ConfApp.DatabaseName
	timeZone := configs.ConfApp.DatabaseTimezone
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, username, password, databaseName, port, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			if configs.ConfApp.AppMode == gin.DebugMode {
				return logger.Default.LogMode(logger.Info)
			}
			return logger.Default.LogMode(logger.Silent)
		}(),
	})

	if err != nil {
		//log.Fatal("cannot connect to database: ", err) // send detailed error to Sentry
		log.Println("cannot connect to database: ", err)
	}

	return db
}

func CreateInstanceDb() *gorm.DB {
	return DbInstance.Instance()
}
