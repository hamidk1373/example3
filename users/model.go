package users

import "gorm.io/gorm"

// User is users model in database
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;uniqueIndex"`
	Phone    string `json:"phone" gorm:"not null;uniqueIndex"`
	Password string `json:"password" gorm:"not null"`
	Address  string `json:"address"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
}

// LoginUser used for login api
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
