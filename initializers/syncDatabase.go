package initializers

import "go-html-test/models"

func SyncDatabase() {
	//миграция модели User
	DB.AutoMigrate(&models.User{})
}
