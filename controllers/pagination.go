package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := c.Request.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(q.Get("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
			// case !math.IsNaN(float64(limit)):
			// 	limit = 1
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

//GeneratePaginationFromRequest ..
// func GeneratePagination(c *gin.Context) models.Pagination {
// 	// Initializing default
// 	//	var mode string
// 	limit := 2
// 	page := 1
// 	sort := "id ASC"
// 	// query := c.Request.URL.Query()
// 	// for key, value := range query {
// 	// 	queryValue := value[len(value)-1]
// 	// 	switch key {
// 	// 	case "limit":
// 	// 		limit, _ = strconv.Atoi(queryValue)
// 	// 	case "page":
// 	// 		page, _ = strconv.Atoi(queryValue)
// 	// 	case "sort":
// 	// 		sort = queryValue
// 	// 	}
// 	// }

// 	if !math.IsNaN(strconv.ParseFloat(c.Query(), 64)) {
// 		limit, _ = strconv.Atoi(c.Query("limit"))
// 	}
// 	if(c.Query("sort") == "1") {

// 	}
// 	return models.Pagination{
// 		Limit: limit,
// 		Page:  page,
// 		Sort:  sort,
// 	}
