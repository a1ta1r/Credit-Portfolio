package main

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
)

func main() {
	godotenv.Load()

	db, err := services.GetConnection()
	if err != nil {
		panic(utils.ConnectionError)
	}

	db.AutoMigrate(&models.User{})

	healthController := controllers.NewHealthController(db)
	userController := controllers.NewUserController(db)
	commonController := controllers.NewCommonController(db)
	paymentPlanController := controllers.NewPaymentPlanController(db)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())
	router.Use(handlers.CorsHandler())

	jwtWrapper := app.NewJwtWrapper(userController)
	jwtMiddleware := jwtWrapper.GetJwtMiddleware()

	secureJWTGroup := router.Group("/")

	secureJWTGroup.Use(jwtMiddleware.MiddlewareFunc())
	{
		//secureJWTGroup.GET("/user", userController.GetUsers)
		secureJWTGroup.GET("/refreshToken", jwtMiddleware.RefreshHandler)
		secureJWTGroup.GET("/health", healthController.HealthCheck)

		secureJWTGroup.GET("/user", userController.GetUsers)
		secureJWTGroup.GET("/user/name/:username", userController.GetUserByName)
		secureJWTGroup.POST("/user/update", userController.UpdateUser)
		//secureJWTGroup.GET("/user/:id", userController.GetUser)
		secureJWTGroup.POST("/user", userController.AddUser)
		secureJWTGroup.DELETE("/user/:id", userController.DeleteUser)
		secureJWTGroup.GET("auth/:token")

		secureJWTGroup.GET("/bank/:id", commonController.GetBank)
		secureJWTGroup.POST("/bank", commonController.AddBank)

		secureJWTGroup.GET("/currency/:id", commonController.GetCurrency)
		secureJWTGroup.POST("/currency", commonController.AddCurrency)

		secureJWTGroup.GET("/role/:id", commonController.GetRole)
		secureJWTGroup.POST("/role", commonController.AddRole)

		secureJWTGroup.GET("/paymentType/:id", commonController.GetPaymentType)
		secureJWTGroup.POST("/paymentType", commonController.AddPaymentType)

		secureJWTGroup.GET("/plan", paymentPlanController.GetPaymentPlans)
		secureJWTGroup.GET("/plan/:id", paymentPlanController.GetPaymentPlan)
		secureJWTGroup.POST("/plan", paymentPlanController.AddPaymentPlan)
		secureJWTGroup.DELETE("/plan/:id", paymentPlanController.DeletePaymentPlan)
	}

	router.POST("/signin", jwtMiddleware.LoginHandler)
	router.POST("/signup", userController.AddUserAnonymous)

	router.NoRoute(controllers.NotFound)

	router.Run()
}
