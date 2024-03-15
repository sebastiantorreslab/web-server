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
	Email    string `json:"email" binding:"required"`
}

type UserPatchRequest struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

var Users = []User{
	{1, "Juan pablo", "email@gmail.com"},
	{2, "José", "email@gmail.com"},
	{3, "Jaime Pogo", "email@gmail.com"},
	{4, "Jenaro Pinto", "email@gmail.com"},
}

func main() {

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, Users)
	})

	r.PATCH("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userReq UserPatchRequest
		err = c.BindJSON(&userReq)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for k, user := range Users {
			if user.ID == userIDInt {
				if userReq.ID != nil {
					user.ID = *userReq.ID
				}
				if userReq.Username != nil {
					user.Username = *userReq.Username
				}
				if userReq.Email != nil {
					user.Email = *userReq.Email
				}
				Users[k] = user
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", userIDInt)})
		return

	})

}
