package main

import (
	"errors"
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	app.LoadConfig()
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if db.DB() == nil {
		panic(errors.New("could not connect to database"))
	}
	db.Table("currency")
	if err != nil {
		panic(err)
	}
	healthController := controllers.NewHealthController(db)
	calculatorController := controllers.NewCalculatorController(db)

	r := gin.Default()
	r.GET("/health", healthController.HealthCheck)
	r.POST("/calculator", calculatorController.Calculate)
	r.NoRoute(controllers.NotFound)
	r.Run()
}
