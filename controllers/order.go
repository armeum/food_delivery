package controllers

// func OrderFood(c *gin.Context) {
// 	//get model if exists
// 	var cart []models.Cart
// 	var product models.Product
// 	db := c.MustGet("db").(*gorm.DB)
// 	if err := db.Where("title = ?", c.Param("title")).Find(&product).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message":    "Route GET:/product/:title not found",
// 			"error":      "Record not found",
// 			"statusCode": 404,
// 		})
// 		return
// 	}

// 	// cart = append(cart.Order, &product)

// 	c.JSON(http.StatusOK, gin.H{"data": product})

// }
