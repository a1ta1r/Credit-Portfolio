package main

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/storages"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := app.GetConnection()
	defer db.Close()
	if err != nil {
		panic(codes.ConnectionError)
	}

	// db.AutoMigrate(
	// 	&models.Bank{},
	// 	&models.Currency{},
	// 	&models.User{},
	// 	&models.PaymentPlan{},
	// 	&models.Payment{},
	// 	&models.Income{},
	// 	&models.Expense{},
	// )

	storageContainer := storages.NewStorageContainer(db)

	//Add services to DI
	userService := services.NewUserService(storageContainer)

	healthController := controllers.NewHealthController(&db)
	userController := controllers.NewUserController(userService)
	commonController := controllers.NewCommonController(&db)
	paymentPlanController := controllers.NewPaymentPlanController(&db, userService, services.PaymentPlanService{})
	paymentController := controllers.NewPaymentController(&db)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())
	router.Use(handlers.CorsHandler())

	jwtWrapper := app.NewJwtWrapper(userService)
	jwtMiddleware := jwtWrapper.GetJwtMiddleware()

	secure := router.Group("/")
	//secure.Use(jwtMiddleware.MiddlewareFunc())
	secure.Use(jwtMiddleware.MiddlewareFunc())
	{
		secure.GET("/refreshToken", jwtMiddleware.RefreshHandler)

		secure.GET("/users", userController.GetUsers)
		secure.GET("/users/:username", userController.GetUserByUsername)
		secure.PUT("/users", userController.UpdateUser)
		secure.DELETE("/users", userController.DeleteUser)
		secure.GET("/user", userController.GetUserByJWT)

		secure.GET("/plan", paymentPlanController.GetPaymentPlans)

		secure.GET("/plan/:id", paymentPlanController.GetPaymentPlan)
		secure.POST("/plan", paymentPlanController.AddPaymentPlan)
		secure.PUT("/plan/:id", paymentPlanController.UpdatePaymentPlan)
		secure.DELETE("/plan/:id", paymentPlanController.DeletePaymentPlan)

		secure.GET("/plan/:id/payments", paymentController.GetPaymentsByPlan)
		secure.GET("/payment/:id", paymentController.GetPayment)
		secure.POST("/payment", paymentController.AddPayment)
		secure.DELETE("/payment/:id", paymentController.DeletePayment)
	}

	router.GET("/health", healthController.HealthCheck)

	router.POST("/signin", jwtMiddleware.LoginHandler)
	router.POST("/signup", userController.AddUser)

	router.GET("/bank/:id", commonController.GetBank)
	router.POST("/bank", commonController.AddBank)
	router.DELETE("/bank/:id", commonController.DeleteBank)
	router.PUT("/bank/:id", commonController.UpdateBank)

	router.GET("/currency/:id", commonController.GetCurrency)
	router.POST("/currency", commonController.AddCurrency)
	router.DELETE("/currency/:id", commonController.DeleteCurrency)
	router.PUT("/currency/:id", commonController.UpdateCurrency)

	router.NoRoute(controllers.NotFound)

	router.Run()
}
