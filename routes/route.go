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

	stationRepository := repositories.NewStationRepository(db)
	stationUsecase := usecases.NewStationUsecase(stationRepository)
	stationController := controllers.NewStationController(stationUsecase)

	api.GET("/admin/stations", stationController.GetAllStations)
	api.GET("/admin/stations/:id", stationController.GetStationByID)
	api.PUT("/admin/stations/:id", stationController.UpdateStation)
	api.POST("/admin/stations", stationController.CreateStation)
	api.DELETE("/admin/stations/:id", stationController.DeleteStation)

	trainRepository := repositories.NewTrainRepository(db)
	trainUsecase := usecases.NewTrainUsecase(trainRepository)
	trainController := controllers.NewTrainController(trainUsecase)

	api.GET("/admin/trains", trainController.GetAllTrains)
	api.GET("/admin/trains/:id", trainController.GetTrainByID)
	api.PUT("/admin/trains/:id", trainController.UpdateTrain)
	api.POST("/admin/trains", trainController.CreateTrain)
	api.DELETE("/admin/trains/:id", trainController.DeleteTrain)

	trainPeronRepository := repositories.NewTrainPeronRepository(db)
	trainPeronUsecase := usecases.NewTrainPeronUsecase(trainPeronRepository)
	trainPeronController := controllers.NewTrainPeronController(trainPeronUsecase)

	api.GET("/admin/train-peron", trainPeronController.GetAllTrainPerons)
	api.GET("/admin/train-peron/:id", trainPeronController.GetTrainPeronByID)
	api.PUT("/admin/train-peron/:id", trainPeronController.UpdateTrainPeron)
	api.POST("/admin/train-peron", trainPeronController.CreateTrainPeron)
	api.DELETE("/admin/train-peron/:id", trainPeronController.DeleteTrainPeron)
}
