package models

type Booking struct {
	Id         int64     `json:"id" gorm:"primary_key"`
	UserId     int64     `json:"userId" gorm:"not null"`
	User       *User     `json:"user" gorm:"foreignKey:UserId;references:Id"`
	PropertyId int64     `json:"propertyId" gorm:"not null"`
	Property   *Property `json:"property" gorm:"foreignKey:PropertyId;references:Id"`
	FullName   string    `json:"fullName" binding:"required"`
	StartDate  string    `json:"startDate" binding:"required"`
	EndDate    string    `json:"endDate" binding:"required"`
	NoOfGuests int64     `json:"noOfGuests"`
	Price      int64     `json:"price"`
	Guidence   string    `json:"guidence"`
	MobileNo   int64     `json:"mobileNo"`
	Email      string    `json:"email"`
}

func (Booking) TableName() string {
	return "bookings"
}
