package routes

import (
	"back-end-golang/controllers"
	"back-end-golang/middlewares"
	"back-end-golang/repositories"
	"back-end-golang/usecases"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func Init(e *echo.Echo, db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// USER

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	api := e.Group("/api/v1")
	api.POST("/login", userController.UserLogin)
	api.POST("/register", userController.UserRegister)

	user := api.Group("/user")
	user.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("user"))
	user.Any("", userController.UserCredential)
	user.PATCH("/update-information", userController.UserUpdateInformation)
	user.PUT("/update-password", userController.UserUpdatePassword)
	user.PUT("/update-profile", userController.UserUpdateProfile)
	user.PUT("/update-photo-profile", userController.UserUpdatePhotoProfile)
	user.DELETE("/delete-photo-profile", userController.UserDeletePhotoProfile)

	historySearchRepository := repositories.NewHistorySearchRepository(db)
	historySearchUsecase := usecases.NewHistorySearchUsecase(historySearchRepository, userRepository)
	historySearchController := controllers.NewHistorySearchController(historySearchUsecase)

	user.GET("/history-search", historySearchController.HistorySearchGetById)
	user.POST("/history-search", historySearchController.HistorySearchCreate)
	user.PUT("/history-search", historySearchController.HistorySearchUpdate)
	user.DELETE("/history-search", historySearchController.HistorySearchDelete)

	// ADMIN

	stationRepository := repositories.NewStationRepository(db)
	stationUsecase := usecases.NewStationUsecase(stationRepository)
	stationController := controllers.NewStationController(stationUsecase)

	trainRepository := repositories.NewTrainRepository(db)
	trainUsecase := usecases.NewTrainUsecase(trainRepository)
	trainController := controllers.NewTrainController(trainUsecase)

	trainPeronRepository := repositories.NewTrainPeronRepository(db)
	trainPeronUsecase := usecases.NewTrainPeronUsecase(trainPeronRepository)
	trainPeronController := controllers.NewTrainPeronController(trainPeronUsecase)

	reservationRepository := repositories.NewReservationRepository(db)
	reservationUsecase := usecases.NewReservationUsecase(reservationRepository)
	reservationController := controllers.NewReservationController(reservationUsecase)

	articleRepository := repositories.NewArticleRepository(db)
	articleUsecase := usecases.NewArticleUsecase(articleRepository)
	articleController := controllers.NewArticleController(articleUsecase)

	recommendationRepository := repositories.NewRecommendationRepository(db)
	recommendationUsecase := usecases.NewRecommendationUsecase(recommendationRepository)
	recommendationController := controllers.NewRecommendationController(recommendationUsecase)

	admin := api.Group("/admin")
	admin.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("admin"))
	admin.GET("/station", stationController.GetAllStations)
	admin.GET("/station/:id", stationController.GetStationByID)
	admin.PUT("/station/:id", stationController.UpdateStation)
	admin.POST("/station", stationController.CreateStation)
	admin.DELETE("/station/:id", stationController.DeleteStation)

	admin.GET("/train", trainController.GetAllTrains)
	admin.GET("/train/:id", trainController.GetTrainByID)
	admin.PUT("/train/:id", trainController.UpdateTrain)
	admin.POST("/train", trainController.CreateTrain)
	admin.DELETE("/train/:id", trainController.DeleteTrain)

	admin.GET("/train-peron", trainPeronController.GetAllTrainPerons)
	admin.GET("/train-peron/:id", trainPeronController.GetTrainPeronByID)
	admin.PUT("/train-peron/:id", trainPeronController.UpdateTrainPeron)
	admin.POST("/train-peron", trainPeronController.CreateTrainPeron)
	admin.DELETE("/train-peron/:id", trainPeronController.DeleteTrainPeron)

	admin.GET("/article", articleController.GetAllArticles)
	admin.GET("/article/:id", articleController.GetArticleByID)
	admin.PUT("/article/:id", articleController.UpdateArticle)
	admin.POST("/article", articleController.CreateArticle)
	admin.DELETE("/article/:id", articleController.DeleteArticle)

	admin.GET("/recommendation", recommendationController.GetAllRecommendations)
	admin.GET("/recommendation/:id", recommendationController.GetRecommendationByID)
	admin.PUT("/recommendation/:id", recommendationController.UpdateRecommendation)
	admin.POST("/recommendation", recommendationController.CreateRecommendation)
	admin.DELETE("/recommendation/:id", recommendationController.DeleteRecommendation)

	api.POST("/reservations", reservationController.AdminCreateReservation)
}
