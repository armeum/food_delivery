package controllers

import (
	"food_delivery/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	// Initializing default
	//	var mode string
	limit := 2
	page := 1
	sort := "category_id ASC"
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
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
