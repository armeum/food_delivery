package controllers

import (
	"encoding/json"
	"fmt"
	"food_delivery/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddItemToBasket(c *gin.Context) {
	//get model if exists
	// db := c.MustGet("db").(*gorm.DB)
	// if err := db.Where("id = ?", c.Param("id")).Find(&item).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route GET:/admin_products/:title not found",
	// 		"error":      "Record not found",
	// 		"statusCode": 404,
	// 	})
	// 	return
	// }
	var item models.Basket
	rBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(rBody, &item)

	item.TotalPrice = item.TotalPrice * float64(item.Unit)

	c.JSON(http.StatusOK, gin.H{"data": item})

}

func DeleteItemFromBasket(c *gin.Context) {
	
}


