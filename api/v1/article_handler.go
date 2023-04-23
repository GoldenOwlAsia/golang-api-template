package v1

import (
	"github.com/GoldenOwlAsia/golang-api-template/api/forms"
	"github.com/GoldenOwlAsia/golang-api-template/api/v1/responses"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"net/http"
)

type ArticleHandler struct {
	DB *gorm.DB
}

func NewArticleHandler(db *gorm.DB) ArticleHandler {
	return ArticleHandler{
		db,
	}
}

func (receiver ArticleHandler) All(c *gin.Context) {
	pagination := responses.Pagination(c)
	baseModel := &models.Article{}
	baseQuery := receiver.DB.Model(baseModel).Preload("User")
	data, paginated, _ := responses.Paginate(baseQuery, baseModel, &pagination)
	c.JSON(http.StatusOK, responses.PaginatedResponse{
		Metadata: paginated,
		Records:  data,
	})
}

func (receiver ArticleHandler) Get(id int64) (*models.Article, error) {
	var article models.Article
	err := receiver.DB.Where("id = ?", id).First(&article).Error
	return &article, err
}

var articleForm = new(forms.ArticleForm)

func (receiver ArticleHandler) Create(c *gin.Context) {
	var form forms.CreateArticleForm
	currentUser := c.MustGet("currentUser").(models.User)
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
	}
	model := models.Article{
		Title:   form.Title,
		Content: form.Content,
		UserID:  currentUser.ID,
	}
	if err := receiver.DB.Create(&model).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article created", "data": model})
}

func (receiver ArticleHandler) Update(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	var form forms.CreateArticleForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
	}
	model := models.Article{ID: id}
	if err := receiver.DB.First(&model).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "model not found!"})
	}
	model.Title = form.Title
	model.Content = form.Content
	if err := receiver.DB.Save(model).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be saved"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article saved", "data": model})
}

func (receiver ArticleHandler) Delete(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	err := receiver.DB.Delete(&models.Article{ID: id}).Error
	c.JSON(http.StatusOK, err)
}
