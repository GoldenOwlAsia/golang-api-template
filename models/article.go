package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

// Article ...
type Article struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
	UserID    uint           `json:"-"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
}

// ArticleModel ...
type ArticleModel struct{}

// All ...
//func (m ArticleModel) All(userID int64) (articles []DataList, err error) {
//	_, err = db.GetDB().Select(&articles, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data, (SELECT row_to_json(n) FROM ( SELECT count(a.id) AS total FROM public.article AS a WHERE a.user_id=$1 LIMIT 1 ) n ) AS meta FROM ( SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC) d", userID)
//	return articles, err
//}

// JSONRaw ...
type JSONRaw json.RawMessage

// DataList ....
type DataList struct {
	Data JSONRaw `db:"data" json:"data"`
	Meta JSONRaw `db:"meta" json:"meta"`
}
