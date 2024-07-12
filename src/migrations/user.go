package migrations

import (
	"Gin/src/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}
