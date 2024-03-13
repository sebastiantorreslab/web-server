package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type product struct {
	Name      string
	Price     int
	Published bool
}

func main() {

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello world",
		})

	})

	router.Run()

	p := product{
		Name:      "Macbook",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	jsonInfo := `{"Name":"Iphone","Price":900,"Published":true}`

	var p1 product

	if err := json.Unmarshal([]byte(jsonInfo), &p1); err != nil {
		log.Fatal(err)
	}

}
