package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateBasketInput struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

type UpdateBasketItemInput struct {
	gorm.Model
	Items []models.BasketItem
}

func UpdateBasket(c *gin.Context) {

	var basket models.Basket
	var basketItems []models.BasketItem
	var total_price uint
	var user_id = uint(c.GetInt("id"))
	// paramInt, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	////checking if user_id exists
	// var basket = find({user_id: c.id})
	db := c.MustGet("db").(*gorm.DB)
	// find basket by user_id_exists

	// if err := db.Where("user_id = ?", user_id).Find(&basket).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route GET:/getAllCategories not found",
	// 		"error":      "Record not found",
	// 		"statusCode": 404,
	// 	})

	// 	return
	// }

	// if basket.UserID != user_id {
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"message":    "Route GET:/getAllCategories not found",
	// 		"error":      "Record not found",
	// 		"statusCode": http.StatusForbidden,
	// 	})
	// 	return
	// }

	if err := db.Where("basket_id = ?", basket.ID).Find(&basketItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/basketitems not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	fmt.Println(basketItems)

	// var body []models.BasketItem
	// var itm = new(models.BasketItem)
	// itm.BasketID = basket.ID
	// itm.ProductID = 3
	// itm.Quantity = 1
	// body = append(body, *itm)
	// fmt.Println(body[0], "item")

	// var res []models.BasketItem
	// for i := 0; i < len(body); i++ {
	// 	// if item.ProductID  {
	// 	if ok := aa[body[i].ProductID]; ok != nil {
	// 		fmt.Println("true")
	// 		fmt.Println(ok, "ok")
	// 		// fmt.Println(val, "val")
	// 		// aa[body[i].ProductID] = val
	// 		total_price += body[i].Quantity * body[i].Product.Price
	// 	} else {
	// 		fmt.Println("false")
	// 		fmt.Println(ok, "ok")
	// 		// fmt.Println(val, "val")
	// 		fmt.Println(body[i], "item")

	// 		aa[body[i].ProductID] = body[i]
	// 		total_price += body[i].Quantity * body[i].Product.Price
	// 	}
	// 	res = append(res, aa[body[i].ID])
	// 	// }
	// }

	// var aa = map[uint]models.BasketItem{
	// 	1: {
	// 		ProductID: 1,
	// 		Quantity: 1,
	// 	},
	// }

	// // err := json.Unmarshal([]byte(aa), &aa)

	// for _, item := range basketItems {
	// 	aa[item.ID] = item
	// }
	var res []models.BasketItem

	body_map := map[uint]models.BasketItem{}

	for p_id, basketItem := range body_map {
		fmt.Println(p_id, "item")
		fmt.Println(basketItem, "basketItem")
		res = append(res, basketItem)
	}

	var body []models.BasketItem = []models.BasketItem{}
	var itm = new(models.BasketItem)
	itm.BasketID = basket.ID
	itm.ProductID = 3
	itm.Quantity = 1
	total_price += 0
	body = append(body, *itm)
	fmt.Println(body, "item")
	fmt.Println(body[0], "body[0]")

	for aaa, item := range body {
		fmt.Println(item, "item")
		fmt.Println(&aaa, "_")

	}

	var aa = map[uint]models.BasketItem{}

	for i := 0; i < len(body); i++ {
		// if item.ProductID  {
		if ok := body[i].ProductID; ok == body[i].ProductID {
			fmt.Println("true")
			fmt.Println(ok, "ok")
			// fmt.Println(val, "val")
			// aa[body[i].ProductID] = ok
			total_price += body[i].Quantity * body[i].Product.Price
		} else {
			fmt.Println("false")
			fmt.Println(ok, "ok")
			// fmt.Println(val, "val")
			fmt.Println(body[i], "item")

			aa[body[i].ProductID] = body[i]
			total_price += body[i].Quantity * body[i].Product.Price
		}
		res = append(res, aa[body[i].ID])
		// }
	}

	fmt.Println(res)
	db.Model(&basketItems).Updates(res)
	fmt.Println(body_map, "body_map")
	fmt.Println(user_id)
	fmt.Println(*itm)
	fmt.Println(c.Get("phone_number"))
	fmt.Println(c.Get("id"))

	c.JSON(http.StatusOK, gin.H{"data": basket})
}
