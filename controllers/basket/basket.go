package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddBasketInput struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

type UpdateBasketInput1 struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

type UpdateBasketItemInput1 struct {
	gorm.Model
	// BasketID  uint `json:"basket_id"`
	Items []models.BasketItem
	// ProductID uint `json:"product_id"`
}

// func GetBasket(c *gin.Context) {
// 	var basket models.Basket
// 	var user_id = c.GetInt("id")
// 	////checking if user_id exists
// 	// if !user_id_exists {
// 	// 	c.JSON(401, gin.H{"message": "user_id not found"})
// 	// }
// 	fmt.Println(user_id)
// 	db := c.MustGet("db").(*gorm.DB)
// 	if err := db.Where("user_id = ?", user_id).Preload("Item").Find(&basket).Error; err != nil {
// 		newBasket := models.Basket{UserID: uint(user_id), TotalPrice: 0}
// 		db.Create(&newBasket)
// 		newBasket.Item = []models.BasketItem{}
// 		c.JSON(http.StatusOK, gin.H{
// 			"message":    "Route GET:/getAllCategories not found",
// 			"error":      "Record not found",
// 			"statusCode": 200,
// 			"data":       newBasket,
// 		})
// 		return
// 		// c.JSON(http.StatusBadRequest, gin.H{
// 		// 	"message":    "Route GET:/getAllCategories not found",
// 		// 	"error":      "Record not found",
// 		// 	"statusCode": 404,
// 		// })
// 		// return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": basket})
// }

// func AddItem(c *gin.Context) {
// 	var basket_item models.BasketItem
// 	var input AddBasketInput
// 	db := c.MustGet("db").(*gorm.DB)

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message":    "Route PUT:/basket not found",
// 			"error":      "Record not found",
// 			"statusCode": 404,
// 		})
// 		return
// 	}

// 	//Create product
// 	basketItem := models.BasketItem{BasketID: basket_item.BasketID, ProductID: basket_item.ProductID, Quantity: basket_item.Quantity}
// 	db.Create(&basketItem)
// 	c.JSON(http.StatusOK, gin.H{"data": basketItem})

// }

func UpdateBasket1(c *gin.Context) {

	// var products models.Product
	var basket models.Basket
	var basketItems []models.BasketItem
	// var user models.User
	var total_price uint
	// var updateBasketItemInput UpdateBasketItemInput

	var user_id = uint(c.GetInt("id"))
	// user_idInt := int(user_id)
	////checking if user_id exists
	// var basket = find({user_id: c.id})
	db := c.MustGet("db").(*gorm.DB)
	//// find basket by user_id_exists
	if err := db.Where("id = ?", user_id).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})

		return
	}
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

	var aa = []models.BasketItem{}

	for _, item := range basketItems {
		aa[item.ID] = item
	}

	fmt.Println(aa)

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

	body_map := (map[uint]models.BasketItem{
		3: {
			BasketID:  basket.ID,
			ProductID: 3,
			Quantity:  1,
		},
	})
	var res []models.BasketItem = []models.BasketItem{}

	for p_id, basketItem := range body_map {
		fmt.Println(p_id, "item")
		fmt.Println(basketItem, "basketItem")
		res = append(res, basketItem)

	}
	return

	var body []models.BasketItem
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
		fmt.Println(aaa, "_")

	}

	// var res []models.BasketItem
	// for i := 0; i < len(body); i++ {
	// 	// if item.ProductID  {
	// 	if val, ok := aa[body[i].ProductID]; ok {
	// 		fmt.Println("true")
	// 		fmt.Println(ok, "ok")
	// 		fmt.Println(val, "val")
	// 		aa[body[i].ProductID] = val
	// 		total_price += body[i].Quantity * body[i].Product.Price
	// 	} else {
	// 		fmt.Println("false")
	// 		fmt.Println(ok, "ok")
	// 		fmt.Println(val, "val")
	// 		fmt.Println(body[i], "item")

	// 		aa[body[i].ProductID] = body[i]
	// 		total_price += body[i].Quantity * body[i].Product.Price
	// 	}
	// 	res = append(res, aa[body[i].ID])
	// 	// }
	// }

	fmt.Println(res)
	db.Model(&basketItems).Updates(res)
	// fmt.Println(user_id)
	// fmt.Println(c.Get("phone_number"))
	// fmt.Println(c.Get("id"))

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetBasketById1(c *gin.Context) {
	var basket []models.Basket
	var user_id = c.GetInt("id")
	// user_idInt := int(user_id)
	////checking if user_id exists

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: uint(user_id), TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"message":    "Created a new basket",
			"error":      "",
			"statusCode": 200,
			"data":       newBasket,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
