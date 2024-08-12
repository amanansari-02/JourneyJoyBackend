package models

type ContactUs struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	FullName string `binding:"required" json:"fullName" bson:"fullName"`
	Email    string `binding:"required" json:"email" bson:"email"`
	Message  string `json:"message" bson:"message"`
}

func (ContactUs) TableName() string {
	return "contact_us"
}
