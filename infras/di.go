package infras

import (
	controllers2 "github.com/GoldenOwlAsia/golang-api-template/handlers"
	"gorm.io/gorm"
)

type AppHandler struct {
	User    controllers2.UserHandler
	Article controllers2.ArticleHandler
}

func DI(db *gorm.DB) AppHandler {
	return AppHandler{
		User:    InitUserAPI(db),
		Article: InitArticleAPI(db),
	}
}
