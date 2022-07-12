package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//Creo las constanstes para las rutas de la api

const (
	PhotoApi = "https://api.pexels.com/v1/"
	VideoApi = "https://api.pexels.com/videos/"
)

func HomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		err.Error()
		fmt.Print(err.Error())
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
		"age":  age,
	})
}
func ParamStrings(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {

	r := gin.Default()
	r.GET("/ping", HomePage)
	r.GET("/query", QueryStrings)
	r.GET("/param/:name/:age", ParamStrings)

	r.Run()
}
