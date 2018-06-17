package main

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := services.GetConnection()
	if err != nil {
		panic(utils.ConnectionError)
	}

	//db.AutoMigrate(
	//	&models.Bank{},
	//	&models.Currency{},
	//	&models.PaymentType{},
	//	&models.Role{},
	//	&models.TimePeriod{},
	//	&models.User{},
	//	&models.PaymentPlan{},
	//	&models.Payment{},
	//	&models.Income{},
	//	&models.Expense{},
	//)

	healthController := controllers.NewHealthController(db)
	userController := controllers.NewUserController(db)
	commonController := controllers.NewCommonController(db)
	paymentPlanController := controllers.NewPaymentPlanController(db)
	paymentController := controllers.NewPaymentController(db)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())
	router.Use(handlers.CorsHandler())

	jwtWrapper := app.NewJwtWrapper(userController)
	jwtMiddleware := jwtWrapper.GetJwtMiddleware()

	secureJWTGroup := router.Group("/")

	secureJWTGroup.Use(jwtMiddleware.MiddlewareFunc())
	{
		secureJWTGroup.GET("/refreshToken", jwtMiddleware.RefreshHandler)

		router.POST("/period", commonController.AddTimePeriod)

		secureJWTGroup.GET("/users", userController.GetUsers)
		secureJWTGroup.GET("/user/name/:username", userController.GetUserByName)
		secureJWTGroup.POST("/user/update", userController.UpdateUser)

		secureJWTGroup.DELETE("/user/:id", userController.DeleteUser)
		secureJWTGroup.GET("auth/:token")

		secureJWTGroup.GET("/plan", paymentPlanController.GetPaymentPlans)
		secureJWTGroup.GET("/plan/:id", paymentPlanController.GetPaymentPlan)
		secureJWTGroup.POST("/plan", paymentPlanController.AddPaymentPlan)
		secureJWTGroup.DELETE("/plan/:id", paymentPlanController.DeletePaymentPlan)

		secureJWTGroup.GET("/plan/:id/payments", paymentController.GetPaymentsByPlan)
		secureJWTGroup.GET("/payment/:id", paymentController.GetPayment)
		secureJWTGroup.POST("/payment", paymentController.AddPayment)
		secureJWTGroup.DELETE("/payment/:id", paymentController.DeletePayment)

		secureJWTGroup.GET("/user", userController.GetUser)
	}

	router.GET("/health", healthController.HealthCheck)

	router.POST("/signin", jwtMiddleware.LoginHandler)
	router.POST("/signup", userController.AddUserAnonymous)

	router.GET("/bank/:id", commonController.GetBank)
	router.POST("/bank", commonController.AddBank)

	router.GET("/currency/:id", commonController.GetCurrency)
	router.POST("/currency", commonController.AddCurrency)

	router.GET("/role/:id", commonController.GetRole)
	router.POST("/role", commonController.AddRole)

	router.GET("/period", commonController.GetAllTimePeriods)
	router.GET("/period/:name", commonController.GetTimePeriod)

	router.GET("/paymentType/:id", commonController.GetPaymentType)
	router.POST("/paymentType", commonController.AddPaymentType)

	router.NoRoute(controllers.NotFound)

	router.Run()
}
