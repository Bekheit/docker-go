package main

import (
	"Docker-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run()
}
