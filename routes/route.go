package routes

import (
	"back-end-golang/controllers"
	"back-end-golang/middlewares"
	"back-end-golang/repositories"
	"back-end-golang/usecases"
	"log"
	"net/http"
	"os"

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

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	stationRepository := repositories.NewStationRepository(db)
	stationUsecase := usecases.NewStationUsecase(stationRepository)
	stationController := controllers.NewStationController(stationUsecase)

	trainStationRepository := repositories.NewTrainStationRepository(db)
	// trainStationUsecase := usecases.NewTrainStationUsecase(trainStationRepository)
	// trainStationController := controllers.NewTrainStationController(trainStationUsecase)

	trainRepository := repositories.NewTrainRepository(db)
	trainUsecase := usecases.NewTrainUsecase(trainRepository, trainStationRepository)
	trainController := controllers.NewTrainController(trainUsecase)

	trainCarriageRepository := repositories.NewTrainCarriageRepository(db)
	trainCarriageUsecase := usecases.NewTrainCarriageUsecase(trainCarriageRepository, trainRepository)
	trainCarriageController := controllers.NewTrainCarriageController(trainCarriageUsecase)

	reservationRepository := repositories.NewReservationRepository(db)
	reservationImageRepository := repositories.NewReservationImageRepository(db)
	reservationUsecase := usecases.NewReservationUsecase(reservationRepository, reservationImageRepository)
	reservationController := controllers.NewReservationController(reservationUsecase, reservationImageRepository)

	articleRepository := repositories.NewArticleRepository(db)
	articleUsecase := usecases.NewArticleUsecase(articleRepository)
	articleController := controllers.NewArticleController(articleUsecase)

	recommendationRepository := repositories.NewRecommendationRepository(db)
	recommendationUsecase := usecases.NewRecommendationUsecase(recommendationRepository)
	recommendationController := controllers.NewRecommendationController(recommendationUsecase)

	historySearchRepository := repositories.NewHistorySearchRepository(db)
	historySearchUsecase := usecases.NewHistorySearchUsecase(historySearchRepository, userRepository)
	historySearchController := controllers.NewHistorySearchController(historySearchUsecase)

	userByAdminRepository := repositories.NewUserRepository(db)
	userByAdminUsecase := usecases.NewUserByAdminUsecase(userByAdminRepository)
	userByAdminController := controllers.NewUserByAdminController(userByAdminUsecase)

	api := e.Group("/api/v1")

	// USER
	api.POST("/login", userController.UserLogin)
	api.POST("/register", userController.UserRegister)

	user := api.Group("/user")
	user.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("user"))

	// user account
	user.Any("", userController.UserCredential)
	user.PATCH("/update-information", userController.UserUpdateInformation)
	user.PUT("/update-password", userController.UserUpdatePassword)
	user.PUT("/update-profile", userController.UserUpdateProfile)
	user.PUT("/update-photo-profile", userController.UserUpdatePhotoProfile)
	user.DELETE("/delete-photo-profile", userController.UserDeletePhotoProfile)

	// train ka
	user.GET("/train/search", trainController.SearchTrainAvailable)

	user.GET("/history-search", historySearchController.HistorySearchGetAll)
	user.POST("/history-search", historySearchController.HistorySearchCreate)
	user.DELETE("/history-search/:id", historySearchController.HistorySearchDelete)

	// ADMIN

	admin := api.Group("/admin")
	admin.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("admin"))

	// crud station
	admin.GET("/station", stationController.GetAllStations)
	admin.GET("/station/:id", stationController.GetStationByID)
	admin.PUT("/station/:id", stationController.UpdateStation)
	admin.POST("/station", stationController.CreateStation)
	admin.DELETE("/station/:id", stationController.DeleteStation)

	// crud train
	admin.GET("/train", trainController.GetAllTrains)
	admin.GET("/train/:id", trainController.GetTrainByID)
	admin.PUT("/train/:id", trainController.UpdateTrain)
	admin.POST("/train", trainController.CreateTrain)
	admin.DELETE("/train/:id", trainController.DeleteTrain)

	admin.GET("/train-carriage", trainCarriageController.GetAllTrainCarriages)
	admin.GET("/train-carriage/:id", trainCarriageController.GetTrainCarriageByID)
	admin.PUT("/train-carriage/:id", trainCarriageController.UpdateTrainCarriage)
	admin.POST("/train-carriage", trainCarriageController.CreateTrainCarriage)
	admin.DELETE("/train-carriage/:id", trainCarriageController.DeleteTrainCarriage)

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

	admin.GET("/user", userByAdminController.GetAllUserByAdmin)
	admin.GET("/user/:id", userByAdminController.GetUserByAdminByID)
	admin.PUT("/user/:id", userByAdminController.UpdateUserByAdmin)
	admin.POST("/user", userByAdminController.CreateUserByAdmin)
	admin.DELETE("/user/:id", userByAdminController.DeleteUserByAdmin)

	api.POST("/reservations", reservationController.AdminCreateReservation)
	admin.GET("/reservations", reservationController.GetAllReservation)
	admin.POST("/reservations", reservationController.AdminCreateReservation)
	admin.GET("/images/:imageName", func(c echo.Context) error {
		imageName := c.Param("imageName")
		imagePath := "./images/" + imageName
		// Check if the image file exists
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Image not found",
			})
		}
		// Return the image file
		return c.File(imagePath)
	})
}
