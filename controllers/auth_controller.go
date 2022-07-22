package controllers

import (
	"bytes"
	"encoding/json"
	"headfirstgo/food_delivery/models"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type LoginBody struct {
	PhoneNumber string `json:"phone_number"`
}

func Login(c *gin.Context) {

	//validate input
	var input LoginBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("phone_number = ?", c.Param("phone_number")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!!"})
		return
	}
	user.Password = RandomPassword()

	SmsSender(user.PhoneNumber, user.Password)
	db.Model(&user).Updates(user)

	c.JSON(http.StatusOK, gin.H{"data": "Success"})

}

func RandomPassword() string {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(9999999)
	return strconv.Itoa(randomNumber)
}

func SmsSender(phone string, password string) {
	base_url := "https://api.telegram.org/bot"

	values := map[string]string{"phone": phone, "password": password}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(base_url, "aplication/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

}
