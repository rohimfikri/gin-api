package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        string
	FirstName string
	LastName  string
	Password  string
	Email     string
	Active    bool
}
