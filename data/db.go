package data

import (
	"fmt"
	userModel "hexagonal-architecture/app/users/domain/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// DB The database connection
	db *gorm.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	// Migrate the schema
	dbCon.AutoMigrate(&userModel.User{})

	// defer db.Close()
	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *gorm.DB {
	return db
}
