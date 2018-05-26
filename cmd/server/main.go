package main

import (
	"github.com/a1ta1r/credit-portfolio/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.Run()
}
