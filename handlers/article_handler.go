package handlers

import (
	"github.com/GoldenOwlAsia/golang-api-template/handlers/requests"
	"github.com/GoldenOwlAsia/golang-api-template/handlers/responses"
	"github.com/GoldenOwlAsia/golang-api-template/services"
	"net/http"

	"github.com/GoldenOwlAsia/golang-api-template/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ArticleHandler struct {
	Service services.ArticleService
}

func NewArticleHandler(service services.ArticleService) ArticleHandler {
	return ArticleHandler{
		Service: service,
	}
}

func (h *ArticleHandler) All(c *gin.Context) {
	pagination := responses.Pagination(c)
	baseQuery := h.Service.All("User")
	data, paginated, err := responses.Paginate(baseQuery, &models.Article{}, &pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, responses.PaginatedResponse{Metadata: paginated, Records: data})
}

func (h *ArticleHandler) Get(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	article, err := h.Service.Get(id)
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
	user := c.MustGet("currentUser").(models.User)
	article, err := h.Service.Create(form, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article created", "data": article})
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	var form requests.CreateArticleForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	article, err := h.Service.Update(id, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article updated", "data": article})
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))
	err := h.Service.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
