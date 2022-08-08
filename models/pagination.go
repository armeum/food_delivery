package models

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

// func GetPaginData(c *gin.Context) Pagination {
//  var	int8 limit
// 	// query.limit num bo'lsa limit = q.limit else 10
// 	return {
// 		Limit:limit,
// 		Page: c.Query.page,
// 		Sort: c.Query.sort,
// 	}
// }
