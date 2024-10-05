package migrations

import (
	"JourneyJoyBackend/src/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Booking{})
	if err != nil {
		return err
	}

	return nil
}
