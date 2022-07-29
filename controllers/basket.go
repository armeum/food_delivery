package controllers

// import (
// 	"food_delivery/models"

// 	"github.com/gin-gonic/gin"
// )

// type CartInterface interface {
// 	Add(item models.Product) *models.Basket
// }

// func (c *models.Basket) AddItemToBasket(item models.Product) *models.Basket {
// 	current_items := b.GetItems()
// 	exists := false
// 	var current_quantity int64
// 	for _, it := range current_items {
// 		if it.GetItemId() == item.GetItemId() {
// 			exists = true
// 			current_quantity = it.GetItemQuantity()
// 			c.Remove(it)
// 			break
// 		}
// 	}

// 	if exists {
// 		new_quantity := current_quantity + item.GetItemQuantity()
// 		item.SetItemQuantity(new_quantity)
// 		//copy(c.Items[location:], c.Items[location+1:])
// 		//c.Items[len(c.Items)-1], c.Items[location] = c.Items[location], c.Items[len(c.Items)-1]
// 		//c.Items = c.Items[:len(c.Items)-1]
// 	} else {
// 		item.SetItemQuantity(item.GetItemQuantity())
// 	}

// 	//new_item := []Item{item}
// 	c.Items = append(c.Items, item)
// 	c.calculateValue()
// 	return c
// }

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
// var item models.Basket
// rBody, err := ioutil.ReadAll(c.Request.Body)
// if err != nil {
// 	fmt.Println(err)
// }
// json.Unmarshal(rBody, &item)

// item.TotalPrice = item.TotalPrice * float64(item.Unit)

// c.JSON(http.StatusOK, gin.H{"data": item})

// func DeleteItemFromBasket(c *gin.Context) {

// }
