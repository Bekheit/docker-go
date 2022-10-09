package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
}

var users = []User{
	{Name: "mohamed"},
	{Name: "ahmed"},
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", GetAll)
	router.Run()
}
