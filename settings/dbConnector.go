package settings

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB returns a pointer to the database
func GetDB() (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	fmt.Println(dbUser)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database conncted:)")
	return db, nil
}
