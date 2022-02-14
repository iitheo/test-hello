package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run()
}

func IndexHandler(c *gin.Context) {
	//name := c.Params.ByName("name")
	//c.JSON(200, gin.H{
	//	"message": "hello world",
	//})
	c.XML(200, Person{FirstName: "Theo", LastName: "K"})
}
