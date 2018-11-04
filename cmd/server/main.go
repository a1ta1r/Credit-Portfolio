package main

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/app"
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	adsControllers "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/auth"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/common"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/finance/controllers"
	loanControllers "github.com/a1ta1r/Credit-Portfolio/internal/components/loans/controllers"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/system"
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/a1ta1r/Credit-Portfolio/internal/docs" //swagger
)

// @title Loan Portfolio API doc
// @version 0.5
// @description Документация по методам API приложения "Кредитный портфель"
func main() {
	godotenv.Load()

	db, err := app.GetConnection()
	defer db.Close()
	if err != nil {
		panic(codes.ConnectionError)
	}

	//println("Dropping all tables")
	//app.DropAllTables()
	//println("Done. Tables dropped.")

	//println("Synchronizing entities with DB")
	//app.SyncModelsWithSchema()
	//println("Done. DB Modified.")

	storageContainer := common.NewStorageContainer(db)

	//Add services to DI
	userService := services.NewUserService(storageContainer)
	agendaService := services.NewAgendaService(db)

	healthController := system.NewHealthController(&db)
	userController := loanControllers.NewUserController(userService)
	commonController := controllers.NewCommonController(&db)
	paymentPlanController := loanControllers.NewPaymentPlanController(&db, userService)
	paymentController := loanControllers.NewPaymentController(&db)
	incomeController := loanControllers.NewIncomeController(&db)
	expenseController := loanControllers.NewExpenseController(&db)
	agendaController := loanControllers.NewAgendaController(agendaService)
	calculationController := loanControllers.NewCalculatorController(&db)
	advertiserController := adsControllers.NewAdvertiserController(
		storageContainer.AdvertiserStorage,
		storageContainer.BannerStorage,
		storageContainer.BannerPlaceStorage)
	advertisementController := adsControllers.NewAdvertisementController(
		storageContainer.AdvertisementStorage, storageContainer.AdvertiserStorage)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())
	router.Use(handlers.CorsHandler())

	jwtWrapper := auth.NewJwtWrapper(userService)
	userJwtMiddleware := jwtWrapper.GetJwtMiddleware(roles.Basic)
	adminJwtMiddleware := jwtWrapper.GetJwtMiddleware(roles.Admin)
	merchantJwtMiddleware := jwtWrapper.GetJwtMiddleware(roles.Ads)

	router.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//basicAccess := router.Group("/")
	basicAccess := router.Group("/", userJwtMiddleware.MiddlewareFunc())
	{
		basicAccess.GET("/refreshToken", userJwtMiddleware.RefreshHandler)

		basicAccess.GET("/users", userController.GetUsers)
		//basicAccess.GET("/users/:username", userController.GetUserByUsername)
		//basicAccess.PUT("/users", userController.UpdateAdvertiser)
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

	//TODO убрать рекламщиков в вип доступ для админа
	advertisers := router.Group("/partners")
	{
		advertisers.GET("/:id", advertiserController.GetAdvertiser)
		advertisers.GET("", advertiserController.GetAdvertisers)
		advertisers.POST("", advertiserController.AddAdvertiser)
		advertisers.PUT("/:id", advertiserController.UpdateAdvertiser)
		advertisers.DELETE("/:id", advertiserController.DeleteAdvertiser)
		advertisers.GET("/:id/promotions", advertisementController.GetAdvertisementsByAdvertiser)

	}

	advertisements := router.Group("/promotions")
	{
		advertisements.GET("/:id", advertisementController.GetAdvertisement)
		advertisements.GET("", advertisementController.GetAdvertisements)
		advertisements.DELETE("/:id", advertisementController.DeleteAdvertisement)
		advertisements.PUT("/:id", advertisementController.UpdateAdvertisement)
		advertisements.POST("", advertisementController.AddAdvertisement)
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

	router.NoRoute(handlers.NotFound)

	router.Run()
}
