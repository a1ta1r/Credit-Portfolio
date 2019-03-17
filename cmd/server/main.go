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
	statControllers "github.com/a1ta1r/Credit-Portfolio/internal/components/user/controllers"
	statServices "github.com/a1ta1r/Credit-Portfolio/internal/components/user/services"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/user/user_handlers"
	_ "github.com/a1ta1r/Credit-Portfolio/internal/docs" //swagger
	"github.com/a1ta1r/Credit-Portfolio/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)

// @title Loan Portfolio API doc
// @version 0.5
// @description Документация по методам API приложения "Кредитный портфель"
func main() {
	godotenv.Load("env.config")

	db, err := app.GetConnection()
	defer db.Close()
	if err != nil {
		panic(codes.ConnectionError)
	}

	//println("Dropping all tables")
	//app.DropAllTables()
	//println("Done. Tables dropped.")
	//
	println("Synchronizing entities with DB")
	app.SyncModelsWithSchema()
	println("Done. DB Modified.")

	storageContainer := common.NewStorageContainer(db)

	//Add services to DI
	userService := services.NewUserService(storageContainer)
	agendaService := services.NewAgendaService(db)
	userStatService := statServices.UserStatisticsService{storageContainer.UserStorage}

	lastSeenHandler := user_handlers.NewLastSeenHandler(userService)

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
		storageContainer.UserStorage,
		storageContainer.AdvertiserStorage,
		storageContainer.BannerStorage,
		storageContainer.BannerPlaceStorage)
	advertisementController := adsControllers.NewAdvertisementController(
		storageContainer.AdvertisementStorage, storageContainer.AdvertiserStorage)
	userStatController := statControllers.NewUserStatisticsController(userStatService)
	bannersController := adsControllers.NewBannersController(storageContainer.BannerStorage)
	bannerPlacesController := adsControllers.NewBannerPlacesController(storageContainer.BannerPlaceStorage)

	router := gin.New()

	router.Use(handlers.PanicHandler)
	router.Use(gin.Logger())
	router.Use(handlers.CorsHandler())

	jwtWrapper := auth.NewJwtWrapper(userService, storageContainer.AdvertiserStorage)
	userJwtMiddleware := jwtWrapper.GetJwtMiddleware(roles.Basic)
	//adminJwtMiddleware := jwtWrapper.GetJwtMiddleware(roles.Admin)
	//merchantJwtMiddleware := jwtWrapper.GetJwtMiddleware(roles.Ads)

	baseRoute := router.Group(os.Getenv("CREDIT_API_PREFIX"))
	baseRoute.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//basicAccess := router.Group("/")
	basicAccess := baseRoute.Group("/", userJwtMiddleware.MiddlewareFunc(), lastSeenHandler.UpdateLastSeen)
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
	advertisers := baseRoute.Group("/partners")
	{
		advertisers.GET("/:id", advertiserController.GetAdvertiser)
		advertisers.GET("", advertiserController.GetAdvertisers)
		advertisers.POST("", advertiserController.AddAdvertiser)
		advertisers.PUT("/:id", advertiserController.UpdateAdvertiser)
		advertisers.DELETE("/:id", advertiserController.DeleteAdvertiser)
		advertisers.GET("/:id/promotions", advertisementController.GetAdvertisementsByAdvertiser)

	}

	advertisements := baseRoute.Group("/promotions")
	{
		advertisements.GET("/:id", advertisementController.GetAdvertisement)
		advertisements.GET("", advertisementController.GetAdvertisements)
		advertisements.DELETE("/:id", advertisementController.DeleteAdvertisement)
		advertisements.PUT("/:id", advertisementController.UpdateAdvertisement)
		advertisements.POST("", advertisementController.AddAdvertisement)
		advertisements.GET("/:id/banners", bannersController.GetBannersByAdvertisementID)
	}

	banners := baseRoute.Group("/banners")
	{
		banners.GET("/:id", bannersController.GetBannerByID)
		banners.DELETE("/:id", bannersController.DeleteBannerByID)
		banners.PUT("/:id", bannersController.UpdateBanner)
		banners.POST("", bannersController.AddBanner)
		banners.PUT("/:id/views", bannersController.IncrViewsForBanner)
		banners.PUT("/:id/clicks", bannersController.IncrClicksForBanner)
	}

	bannerPlaces := baseRoute.Group("/banner_places")
	{
		bannerPlaces.GET("", bannerPlacesController.GetBannerPlaces)
		bannerPlaces.GET("/:id", bannerPlacesController.GetBannerPlaceByID)
		bannerPlaces.DELETE("/:id", bannerPlacesController.DeleteBannerPlaceByID)
		bannerPlaces.POST("", bannerPlacesController.AddBannerPlace)
		bannerPlaces.PUT("/:id", bannerPlacesController.UpdateBannerPlace)
	}

	systemStat := baseRoute.Group("/stats")
	{
		systemStat.GET("/users/registered", userStatController.GetRegisteredUsersCount)
		systemStat.GET("/users/deleted", userStatController.GetDeletedUsersCount)
		systemStat.GET("/users/active", userStatController.GetLastSeenUsersCount)
	}

	baseRoute.GET("/health", healthController.HealthCheck)

	baseRoute.POST("/signin", userJwtMiddleware.LoginHandler)
	baseRoute.POST("/signup", userController.AddUser)

	router.GET("/bank/:id", commonController.GetBank)
	router.POST("/bank", commonController.AddBank)
	router.DELETE("/bank/:id", commonController.DeleteBank)
	router.PUT("/bank/:id", commonController.UpdateBank)

	router.GET("/currency/:id", commonController.GetCurrency)
	router.POST("/currency", commonController.AddCurrency)
	router.DELETE("/currency/:id", commonController.DeleteCurrency)
	router.PUT("/currency/:id", commonController.UpdateCurrency)

	router.NoRoute(handlers.NotFound)

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	} else if os.Getenv("CREDIT_API_PORT") != "" {
		port = os.Getenv("CREDIT_API_PORT")
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	host := ""
	if os.Getenv("CREDIT_API_HOST") != "" {
		host = os.Getenv("CREDIT_API_HOST")
	}

	router.Run(host + ":" + port)
}
