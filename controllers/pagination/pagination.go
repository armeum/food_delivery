package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := c.Request.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		limit, _ := strconv.Atoi(q.Get("limit"))
		emptyStringToInt, _ := strconv.Atoi("")
		offset := limit * (page - 1)

		if limit == emptyStringToInt {
			limit = 10
		}
		// else if !math.IsNaN(float64(limit)) {
		// 	limit = 10
		// }

		if page == emptyStringToInt {
			page = 1
		} 
		
		fmt.Println(limit, offset)
		return db.Limit(limit).Offset(offset)
	}
}


// func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		q := c.Request.URL.Query()
// 		page, _ := strconv.Atoi(q.Get("page"))
// 		if page == 0 {
// 			page = 1
// 		}

// 		limit := 0

// 		limit, _ = strconv.Atoi(q.Get("limit"))
// 		switch {
// 		case limit >= 100:
// 			limit = 100
// 		case limit == 1:
// 			limit = 1
// 		case limit == 2:
// 			limit = 2
// 		case limit == 3:
// 			limit = 3
// 		case limit == 4:
// 			limit = 4
// 		default:
// 			limit = 10
// 		}
// 		offset := (page - 1) * limit
// 		return db.Offset(offset).Limit(limit)
// 	}
// }

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
