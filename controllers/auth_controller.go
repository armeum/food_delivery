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
	if err := db.Where("phone_number = ?", input.PhoneNumber).First(&user).Error; err != nil {
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
	test(phone, password)
	base_url := "https://api.telegram.org/bot"

	values := map[string]string{"phone": phone, "password": password}
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

func test(phone string, password string) {
	num := strconv.Itoa(-1001685855235)
	httpposturl := fmt.Sprintf("https://api.telegram.org/bot/sendMessage?chat_id=%s&text=Hello+World", num)



	// I think this is the block I need to alter?:
	
	// body := fmt.Sprintf(
	// 	`{
    //     ""chat_id": %d ".}`, num)
		
	// body +=  `"text": "some text"` 
	// body += `phone": "09090990988"`
	var jsonData = []byte(`{
        "text": phone,
        "job": "leader"
    }`)

	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	

	fmt.Println("response Status:", response.Status)

	defer response.Body.Close()
}
