package RadioReader

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func CreateConnection() (*gorm.DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		))

	return db, err
}
