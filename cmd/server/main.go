package main

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/storages"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
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
	agendaService := services.NewAgendaService(db)

	healthController := controllers.NewHealthController(&db)
	userController := controllers.NewUserController(userService)
	commonController := controllers.NewCommonController(&db)
	paymentPlanController := controllers.NewPaymentPlanController(&db, userService, services.PaymentPlanService{})
	paymentController := controllers.NewPaymentController(&db)
	incomeController := controllers.NewIncomeController(&db)
	expenseController := controllers.NewExpenseController(&db)
	agendaController := controllers.NewAgendaController(agendaService)
	calculationController := controllers.NewCalculatorController(&db)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())
	router.Use(handlers.CorsHandler())

	jwtWrapper := app.NewJwtWrapper(userService)
	userJwtMiddleware := jwtWrapper.GetJwtMiddleware(models.Basic)
	adminJwtMiddleware := jwtWrapper.GetJwtMiddleware(models.Admin)
	merchantJwtMiddleware := jwtWrapper.GetJwtMiddleware(models.Ads)

	basicAccess := router.Group("/", userJwtMiddleware.MiddlewareFunc())
	{
		basicAccess.GET("/refreshToken", userJwtMiddleware.RefreshHandler)

		//basicAccess.GET("/users", userController.GetUsers)
		//basicAccess.GET("/users/:username", userController.GetUserByUsername)
		//basicAccess.PUT("/users", userController.UpdateUser)
		//basicAccess.DELETE("/users", userController.DeleteUser)
		//Вроде как юзер не должен напрямую это уметь.
		basicAccess.GET("/user", userController.GetUserByJWT)
		basicAccess.PUT("/user", userController.UpdateUserByJWT)

		basicAccess.GET("/plans", paymentPlanController.GetPaymentPlans)
		basicAccess.GET("/plans/:id", paymentPlanController.GetPaymentPlan)
		basicAccess.POST("/plans", paymentPlanController.AddPaymentPlan)
		basicAccess.PUT("/plans/:id", paymentPlanController.UpdatePaymentPlan)
		basicAccess.DELETE("/plans/:id", paymentPlanController.DeletePaymentPlan)

		basicAccess.GET("/plans/:id/payments", paymentController.GetPaymentsByPlan)
		//basicAccess.GET("/payments/:id", paymentController.GetPayment)
		//basicAccess.POST("/payments", paymentController.AddPayment)
		//basicAccess.DELETE("/payments/:id", paymentController.DeletePayment)
		//Вроде как юзер не должен напрямую это уметь. Сущность агрегируется в пейментплане.

		basicAccess.GET("/incomes", incomeController.GetIncomesByJWT)
		basicAccess.GET("/incomes/:id", incomeController.GetIncomeById)
		basicAccess.PUT("/incomes/:id", incomeController.UpdateIncomeByIdAndJWT)
		basicAccess.POST("/incomes", incomeController.AddIncome)
		basicAccess.DELETE("/incomes/:id", incomeController.DeleteIncomeByIdAndJWT)

		basicAccess.GET("/expenses", expenseController.GetExpensesByJWT)
		basicAccess.GET("/expenses/:id", expenseController.GetExpenseById)
		basicAccess.PUT("/expenses/:id", expenseController.UpdateExpenseByIdAndJWT)
		basicAccess.POST("/expenses", expenseController.AddExpense)
		basicAccess.DELETE("/expenses/:id", expenseController.DeleteExpenseByIdAndJWT)

		basicAccess.POST("/calculate", calculationController.CalculateCredit)

		basicAccess.GET("/agenda", agendaController.GetAgendaElements)
	}

	private := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	}

	//Не работает чет, доступ есть у всего на свете
	adminAccess := router.Group("/administration", adminJwtMiddleware.MiddlewareFunc())
	{
		adminAccess.GET("/private", private)
	}

	//Не работает чет, доступ есть у всего на свете
	merchantAccess := router.Group("/banking", merchantJwtMiddleware.MiddlewareFunc())
	{
		merchantAccess.GET("", private)
	}

	router.GET("/health", healthController.HealthCheck)

	router.POST("/signin", userJwtMiddleware.LoginHandler)
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
