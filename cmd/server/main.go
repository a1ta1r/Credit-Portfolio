package main

import (
	"github.com/gin-gonic/gin"
	"github.com/credit-portfolio/internal/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.Run()
}
