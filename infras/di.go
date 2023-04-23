package infras

import (
	"github.com/GoldenOwlAsia/golang-api-template/handlers"
	"gorm.io/gorm"
)

type AppHandler struct {
	User    handlers.UserHandler
	Article handlers.ArticleHandler
}

func DI(db *gorm.DB) AppHandler {
	return AppHandler{
		User:    InitUserAPI(db),
		Article: InitArticleAPI(db),
	}
}
