package common

type LoginStruct struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"omitempty"`
	IsEmailLogin int64  `json:"isEmailLogin" binding:"omitempty"`
}
