package main

import (
	"github.com/a1ta1r/credit-portfolio/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"errors"
)

func main() {
	godotenv.Load()
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if db.DB() == nil {
		panic(errors.New("could not connect to database"))
	}
	if err != nil {
		panic(err)
	}
	healthController := controllers.NewHealthController(db)
	calculatorController := controllers.NewCalculatorController()

	r := gin.Default()
	r.GET("/health", healthController.HealthCheck)
	r.POST("/calculator", calculatorController.Calculate)
	r.Run()
}
