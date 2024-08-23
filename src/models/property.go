package models

import (
	"time"

	"github.com/lib/pq"
)

type Property struct {
	Id             int64          `json:"id" gorm:"primaryKey"`
	PropertyName   string         `json:"propertyName" bson:"propertyName" binding:"required"`
	PropertyType   string         `json:"propertyType" bson:"propertyType" binding:"required"`
	Price          float64        `json:"price" bson:"price" binding:"required"`
	Description    string         `json:"description" bson:"description"`
	Location       string         `json:"location" bson:"location" binding:"required"`
	City           string         `json:"city" bson:"city" binding:"required"`
	Rooms          int64          `json:"rooms" bson:"rooms" binding:"required"`
	NoOfGuests     int64          `json:"noOfGuests" bson:"noOfGuests" binding:"required"`
	PropertyImages pq.StringArray `json:"propertyImages" gorm:"type:text[]" bson:"propertyImages" binding:"required"`
	CreatedAt      time.Time      `json:"createdAt" bson:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updatedAt" bson:"updatedAt" gorm:"autoUpdateTime"`
}

func (Property) TableName() string {
	return "property"
}
