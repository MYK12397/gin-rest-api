package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Status OK",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body

	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}
func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func main() {

	r := gin.Default()

	r.GET("/", HomePage)
	r.POST("/Post", PostHomePage)

	r.GET("/query", QueryString)              //query?name="yahiya"&age=24
	r.GET("/path/:name/:age", PathParameters) //path/yahiya/24
	r.Run()
}
