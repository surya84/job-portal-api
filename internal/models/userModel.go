package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId       uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
