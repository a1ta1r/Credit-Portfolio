package main

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.LoadConfig()

	db, err := services.GetConnection()
	if err != nil {
		panic(utils.ConnectionError)
	}
	healthController := controllers.NewHealthController(db)
	userController := controllers.NewUserController(db)
	commonController := controllers.NewCommonController(db)

	r := gin.Default()

	r.GET("/health", healthController.HealthCheck)

	r.GET("/user/:id", userController.GetUser)
	r.POST("/user", userController.AddUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	r.GET("/bank/:id", commonController.GetBank)
	r.POST("/bank", commonController.AddBank)

	r.GET("/currency/:id", commonController.GetCurrency)
	r.POST("/currency", commonController.AddCurrency)

	r.GET("/role/:id", commonController.GetRole)
	r.POST("/role", commonController.AddRole)

	r.GET("/paymentType/:id", commonController.GetPaymentType)
	r.POST("/paymentType", commonController.AddPaymentType)

	r.NoRoute(controllers.NotFound)

	r.Run()
}
