package v1

import (
	"net/http"

	"github.com/GoldenOwlAsia/golang-api-template/api/v1/requests"
	"github.com/GoldenOwlAsia/golang-api-template/api/v1/responses"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type ArticleHandler struct {
	DB *gorm.DB
}

func NewArticleHandler(db *gorm.DB) ArticleHandler {
	return ArticleHandler{
		DB: db,
	}
}

func (h *ArticleHandler) All(c *gin.Context) {
	pagination := responses.Pagination(c)
	baseQuery := h.DB.Model(&models.Article{}).Preload("User")
	data, paginated, err := responses.Paginate(baseQuery, &models.Article{}, &pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"metadata": paginated,
		"records":  data,
	})
}

func (h *ArticleHandler) Get(c *gin.Context) {
	id := cast.ToInt64(c.Param("id"))
	var article models.Article
	err := h.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var form requests.CreateArticleForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	currentUser := c.MustGet("currentUser").(models.User)
	article := models.Article{
		Title:   form.Title,
		Content: form.Content,
		UserID:  currentUser.ID,
	}
	if err := h.DB.Create(&article).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article created", "data": article})
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id := cast.ToInt64(c.Param("id"))
	var form requests.CreateArticleForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	article := models.Article{}
	if err := h.DB.First(&article, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	article.Title = form.Title
	article.Content = form.Content
	if err := h.DB.Save(&article).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article updated", "data": article})
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id := cast.ToInt64(c.Param("id"))
	if err := h.DB.Delete(&models.Article{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
