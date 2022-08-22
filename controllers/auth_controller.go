package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	FirstName   string `json:"first_name"`
	PhoneNumber string `gorm:"typevarchar(5);unique_index" json:"phone_number"`
}

 func Login(c *gin.Context) {
	//validate input
	var input LoginBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	//user nil -> userni create qilasz, phone number, '', '', '',
	if err := db.
		Where("phone_number = ?", input.PhoneNumber).
		First(&user); err != nil {
		user := models.User{
			PhoneNumber: input.PhoneNumber,
		}
		db := c.MustGet("db").(*gorm.DB)
		db.Create(&user)
		user.Password = RandomPassword()
		SmsSender(user.FirstName, user.PhoneNumber, user.Password)
		c.JSON(http.StatusOK, gin.H{
			"Password successfully sent to the phone number": user.PhoneNumber,
		})
		db.Save(&user)
	}

	db.Model(&user).Updates(&user)
}

// / Generating random four-digit password
func RandomPassword() string {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(9999-1000) + 1000
	return strconv.Itoa(randomNumber)
}

func SmsSender(first_name string, phone string, password string) {
	test(first_name, phone, password)
	base_url := "https://api.telegram.org/bot"

	values := map[string]string{
		"first_name": first_name,
		"phone":      phone,
		"password":   password,
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(base_url, "aplication/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err, "err")
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["json"], "res")
}

func test(first_name string, phone string, password string) {
	num := strconv.Itoa(-1001685855235)
	httpposturl := fmt.Sprintf("https://api.telegram.org/bot5497289382:AAEAuBV4_JOoU1qwIo9RPktV9X1l7FfOG7o/sendMessage?chat_id=%s&text=%s+%s+%s", num, first_name, phone, password)
	var jsonData = []byte(`{
        "text": phone,
        "job": "leader"
    }`)
	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	fmt.Println("response Status:", response.Status)
	defer response.Body.Close()
}
