package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}
