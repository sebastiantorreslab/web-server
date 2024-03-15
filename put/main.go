package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
}

var Users = []User{
	{1, "Juan pablo"},
	{2, "José"},
	{3, "Jaime Pogo"},
	{4, "Jenaro Pinto"},
}

func main() {

	r := gin.Default()

	r.PUT("/users/:id", func(c *gin.Context) {

		var user User
		c.BindJSON(&user)

		idReceived := c.Param("id")
		idInt, err := strconv.Atoi(idReceived)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for k, u := range Users {
			if u.ID == idInt {
				user.ID = idInt
				Users[k] = user
				c.JSON(http.StatusOK, "Succesfull")
			}
		}

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", idInt)})

	})

	r.Run(":8080")
}
