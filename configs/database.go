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

func DBSeeder(db *gorm.DB) error {
	classTypes := []string{"Ekonomi", "Business", "Eksekutif"}

	for _, class := range classTypes {
		prefixes := []string{"A", "B", "C", "D"}

		if class == "Ekonomi" || class == "Business" {
			prefixes = append(prefixes, "E")
		}

		for _, prefix := range prefixes {
			startIndex := 1

			if prefix == "D" && class == "Ekonomi" || prefix == "D" && class == "Business" || prefix == "E" && class == "Ekonomi" || prefix == "E" && class == "Business" {
				startIndex = 3
			}

			for i := startIndex; i <= 12; i++ {
				seat := models.TrainSeat{
					Class: class,
					Name:  fmt.Sprintf("%s%d", prefix, i),
				}

				// Check if data already exists
				var count int64
				if err := db.Model(&models.TrainSeat{}).Where(&seat).Count(&count).Error; err != nil {
					return err
				}

				// If data exists, skip seeding
				if count > 0 {
					continue
				}

				if class == "Eksekutif" {
					if i >= 1 && i <= 12 {
						seat.Name = fmt.Sprintf("%s%d", prefix, i)
					} else {
						continue
					}
				}

				if err := db.Create(&seat).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Station{},
		&models.Train{},
		&models.TrainPeron{},
		&models.TrainSeat{},
		models.ReservationCheckouts{},
		models.ReservationImages{},
		models.Reservations{},
		models.ReservationImages{},
		&models.Article{},
	)
}
