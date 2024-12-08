package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `gorm:"type:uuid;primary_key;unique" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
