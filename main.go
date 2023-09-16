package main

import (
	"fmt"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Products []string `json:"products"`
}

func main() {
	r := gin.Default()
	r.POST("/sendSms", func(c *gin.Context) {
		var requestBody RequestBody

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email := requestBody.Email
		products := requestBody.Products
		phone := requestBody.Phone

		cartSize := len(products)
		body := fmt.Sprintf("Ei, %v !!\n\nAchou que não iriamos ver??\n\nParece que algo o impediu de continuar com seu carrinho com %v produtos, \n\nNão deixe que essa oportunidade única escape. Seu carrinho de compras está esperando por você:\n\n%v\n\nCom carinho,\n[Tigers]\n\n", email, cartSize, products)

		client := twilio.NewRestClient()

		params := &openapi.CreateMessageParams{}
		params.SetTo(phone)
		params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
		params.SetBody(body)

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("SMS sent successfully!")
		}
	})

	r.POST("/sendWhats", func(c *gin.Context) {
		var requestBody RequestBody

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		email := requestBody.Email
		products := requestBody.Products
		phone := requestBody.Phone

		cartSize := len(products)
		body := fmt.Sprintf("Ei, %v !!\n\nAchou que não iriamos ver??\n\nParece que algo o impediu de continuar com seu carrinho com %v produtos, \n\nNão deixe que essa oportunidade única escape. Seu carrinho de compras está esperando por você:\n\n%v\n\nCom carinho,\n[Tigers]\n\n", email, cartSize, products)

		client := twilio.NewRestClient()

		params := &openapi.CreateMessageParams{}
		to := fmt.Sprintf("whatsapp:%v", phone)
		params.SetTo(to)

		params.SetFrom("whatsapp:+14155238886")
		params.SetBody(body)

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Message sent successfully!")
		}

	})

	r.Run()
}
