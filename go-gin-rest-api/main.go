package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func postHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PathParamStrings(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", homePage)
	r.POST("/", postHomePage)
	r.GET("/query", QueryStrings)
	r.GET("/path/:name/:age", PathParamStrings)

	r.Run()
}