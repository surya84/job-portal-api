package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//UserId       uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Dob          string `json:"dob"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Dob      string `json:"dob" validate:"required"`
}

type UserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Dob   string `json:"dob" validate:"required"`
}

type CheckOtp struct {
	Email           string `json:"email" validate:"required,email"`
	Otp             string `json:"otp" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
