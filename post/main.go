package main

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required`
}

func main() {

	r := gin.Default()

	r.POST("/login", func(c *gin.Context){
		var login Login
		c.BindJSON(&login)
		c.JSON(200,gin.H{
			"status":login.Email,
		})
	})

	r.Run(":8080")


}
