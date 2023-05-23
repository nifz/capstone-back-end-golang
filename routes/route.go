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
	reservationImageRepository := repositories.NewReservationImageRepository(db)
	reservationUsecase := usecases.NewReservationUsecase(reservationRepository, reservationImageRepository)
	reservationController := controllers.NewReservationController(reservationUsecase, reservationImageRepository)

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
