package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb() {

	// Re-using json_crud_api DB
	dsn := "host=localhost user=postgres dbname=json_crud_api port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error opening the database.")
	}

}
