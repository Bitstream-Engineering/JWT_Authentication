package initializers

import "JWT_Authentication/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
