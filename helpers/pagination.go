package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func (pg Pagination) Query(tx *gorm.DB) *gorm.DB {

	offset := (pg.Page - 1) * pg.Limit

	return tx.
		Limit(pg.Limit).
		Offset(offset).
		Order(pg.Sort)
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {

	limit, page := 2, 1
	sort, query := "created_at asc", c.Request.URL.Query()

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

	return Pagination{Limit: limit, Page: page, Sort: sort}
}
