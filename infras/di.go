package infras

import (
	"api/api"
	"api/sse"
	"gorm.io/gorm"
)

func DI(db *gorm.DB, event *sse.Event) api.AppHandler {
	return api.AppHandler{
		User: InitUserAPI(db),
	}
}
