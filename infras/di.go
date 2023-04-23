package infras

import (
	v1 "github.com/GoldenOwlAsia/golang-api-template/api/v1"
	"gorm.io/gorm"
)

type AppHandler struct {
	User    v1.UserHandler
	Article v1.ArticleHandler
}

func DI(db *gorm.DB) AppHandler {
	return AppHandler{
		User:    InitUserAPI(db),
		Article: InitArticleAPI(db),
	}
}
