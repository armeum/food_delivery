package controllers

import (
	"food_delivery/models"
	"food_delivery/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type FavProd struct {
	ProductID uint `gorm:"foreignKey:id" json:"product_id" binding:"required"`
}

type Favorites struct {
	FavProd []*FavProd `json:"fav_items" binding:"required"`
}

func GetFavorites(c *gin.Context) {
	var favorites models.Favorites
	var fav_items []*models.FavItems
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", pkg.GetUserID(c)).Preload("FavItems.Product.Prices.ProductPastry").Find(&favorites).Error; err != nil {
		newFavorites := models.Favorites{UserID: pkg.GetUserID(c)}
		db.Create(&newFavorites)
		newFavorites.FavItems = []*models.FavItems{}
		c.JSON(http.StatusOK, gin.H{
			"data": newFavorites,
		})
		return
	}

	if err := db.Where("favorites_id = ?", favorites.ID).Preload("Product").Find(&fav_items).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": favorites,
	})
}

func AddFavProd(c *gin.Context) {

	var input Favorites

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	favorites, err := getUserFavorites(pkg.GetUserID(c), db)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if favorites.ID == 0 {
		favorites.UserID = pkg.GetUserID(c)
		err := db.Create(&favorites).Error
		if err != nil {
			log.Println(err)
		}
	}

	for _, favs := range input.FavProd {
		var product models.Product

		if err := db.Where("id = ?", favs.ProductID).Find(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	favorites.FavItems = makeUserFavorites(input.FavProd)
	db.Where("favorites_id = ?", favorites.ID).Delete(models.FavItems{})
	db.Save(&favorites)

	c.JSON(http.StatusOK, gin.H{
		"message": favorites,
	})
}

func getUserFavorites(userId uint, db *gorm.DB) (*models.Favorites, error) {
	var favs models.Favorites
	err := db.Where("user_id = ?", userId).Find(&favs).Error
	return &favs, err
}

func makeUserFavorites(favs []*FavProd) []*models.FavItems {

	var userFavorites []*models.FavItems = make([]*models.FavItems, 0)

	for _, fav := range favs {
		userFavs := models.FavItems{}

		userFavs.ProductID = fav.ProductID

		userFavorites = append(userFavorites, &userFavs)
	}
	return userFavorites
}
