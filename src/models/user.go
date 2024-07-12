package models

type User struct {
	Id           int64  `json:"id" gorm:"primary_key`
	Name         string `binding:"required" bson:"name" json:"name"`
	Email        string `binding:"required,email" bson:"email" json:"email"`
	Password     string `binding:"required" bson:"password" json:"password"`
	PhoneNo      string `binding:"omitempty" bson:"phoneNo,omitempty" json:"phoneNo,omitempty"`
	ProfilePhoto string `binding:"omitempty" bson:"profilePhoto" json:"profilePhoto"`
}

func (User) TableName() string {
	return "users"
}
