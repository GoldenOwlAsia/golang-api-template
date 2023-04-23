package services

import (
	"github.com/GoldenOwlAsia/golang-api-template/handlers/requests"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func NewArticleService(db *gorm.DB) ArticleService {
	return ArticleService{
		DB: db,
	}
}

func (s *ArticleService) All(preloads ...string) (baseQuery *gorm.DB) {
	baseQuery = s.DB.Model(&models.Article{})
	for _, preload := range preloads {
		baseQuery = s.DB.Model(&models.Article{}).Preload(preload)
	}
	return
}

func (s *ArticleService) Create(req requests.CreateArticleForm, user models.User) (article *models.Article, err error) {
	article = &models.Article{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user.ID,
	}
	err = s.DB.Create(article).Error
	return
}

func (s *ArticleService) Update(id uint, req requests.CreateArticleForm) (article *models.Article, err error) {
	article = &models.Article{}
	err = s.DB.First(article, id).Error
	if err != nil {
		return
	}
	article.Title = req.Title
	article.Content = req.Content
	err = s.DB.Save(&article).Error
	return
}

func (s *ArticleService) Delete(id uint) (err error) {
	err = s.DB.Delete(&models.Article{}, id).Error
	return
}

func (s *ArticleService) Get(id uint) (article models.Article, err error) {
	err = s.DB.Where("id = ?", id).First(&article).Error
	return
}
