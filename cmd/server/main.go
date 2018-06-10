package main

import (
	"github.com/A1ta1r/Credit-Portfolio/internal/app"
	"github.com/A1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/A1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/A1ta1r/Credit-Portfolio/internal/services"
	"github.com/A1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/appleboy/gin-jwt.v2"
	"time"
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
	paymentPlanController := controllers.NewPaymentPlanController(db)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())

	jwtMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "robreid.io",
		Key:           []byte("portfolio-on-credit-very-very-very-secret-key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: handlers.Authenticate,
		PayloadFunc:  handlers.Payload,
	}

	secureJWTGroup := router.Group("/")

	secureJWTGroup.Use(jwtMiddleware.MiddlewareFunc())
	{
		//secureJWTGroup.GET("/user", userController.GetUsers)
		secureJWTGroup.GET("/refreshToken", jwtMiddleware.RefreshHandler)
		secureJWTGroup.GET("/health", healthController.HealthCheck)

		secureJWTGroup.GET("/user", userController.GetUsers)
		secureJWTGroup.GET("/user/:id", userController.GetUser)
		secureJWTGroup.POST("/user", userController.AddUser)
		secureJWTGroup.DELETE("/user/:id", userController.DeleteUser)

		secureJWTGroup.GET("/hello", handlers.Hello)

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

	router.POST("/login", jwtMiddleware.LoginHandler)

	router.NoRoute(controllers.NotFound)

	router.Run()
}
