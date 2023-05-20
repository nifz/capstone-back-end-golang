package configs

import (
	"back-end-golang/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load the Asia/Jakarta location
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		// Handle the error
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbConn = dbConn.Session(&gorm.Session{
		NowFunc: func() time.Time {
			return time.Now().In(location)
		},
	})

	return dbConn, nil
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
		models.Station{},
		models.Train{},
		models.TrainPeron{},
	)
}
