package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParsePageLimit(c *gin.Context) (page, limit int) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limitStr := c.Query("limit")
	limit, err = strconv.Atoi(limitStr)

	if err != nil {
		limit = 10
	}

	return
}

func ParseSort(s string, queryInput *gorm.DB) (query *gorm.DB) {
	sortSlice := strings.Split(s, ",")
	query = queryInput
	regxOrder, _ := regexp.Compile("^-[a-z]+$")
	for _, sort := range sortSlice {
		m := regxOrder.MatchString(sort)
		if m {
			sort = strings.ReplaceAll(sort, "-", "")
			query = query.Order(fmt.Sprint(sort) + " desc")
		} else {
			query = query.Order(fmt.Sprint(sort) + " asc")
		}
	}

	return
}
