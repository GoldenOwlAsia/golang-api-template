package pkgs

import (
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
)

type Paginate struct {
	Limit      int
	Page       int
	TotalRows  int64
	TotalPages int
}

func NewGinPagy(query gorm.DB, c *gin.Context) *Paginate {
	page, limit := utils.ParsePageLimit(c)
	var totalRows int64
	query.Count(&totalRows)
	return &Paginate{
		TotalRows:  totalRows,
		Limit:      limit,
		Page:       page,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}

func NewPagy(query gorm.DB, c *gin.Context, limit int, page int) *Paginate {
	var totalRows int64
	query.Count(&totalRows)
	return &Paginate{
		TotalRows:  totalRows,
		Limit:      limit,
		Page:       page,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}

func (p *Paginate) Paginated(query *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	return query.Offset(offset).Limit(p.Limit)
}
