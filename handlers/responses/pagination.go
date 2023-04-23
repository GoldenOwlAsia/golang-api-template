package responses

import (
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"reflect"
	"strconv"
)

type PaginatedResponse struct {
	Metadata PaginatedData `json:"_metadata"`
	Records  any           `json:"records"`
}

type PaginatedData struct {
	Limit      int    `json:"limit"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	PerPage    int    `json:"per_page"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
}

func Pagination(c *gin.Context) PaginatedData {
	limit := configs.DefaultItemPerPage
	page := configs.DefaultPage
	sort := configs.DefaultSorting
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}

	return PaginatedData{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func Paginate(baseQuery *gorm.DB, model interface{}, pagination *PaginatedData) (interface{}, PaginatedData, error) {
	offset := (pagination.Page - 1) * pagination.Limit
	perPage := 10

	// Get total count of records
	var count int64
	if err := baseQuery.Count(&count).Error; err != nil {
		return nil, PaginatedData{}, err
	}

	// Calculate pagination values
	totalPages := int64(math.Ceil(float64(count) / float64(perPage)))

	// Build query
	query := baseQuery.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// Retrieve records
	records := reflect.New(reflect.SliceOf(reflect.TypeOf(model).Elem())).Interface()
	if err := query.Find(records).Error; err != nil {
		return nil, PaginatedData{}, err
	}

	return records, PaginatedData{
		Total:      int(count),
		TotalPages: int(totalPages),
		PerPage:    perPage,
		Page:       pagination.Page,
		Sort:       pagination.Sort,
		Limit:      pagination.Limit,
	}, nil
}
