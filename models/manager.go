package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model

	Id    int64
	Name  string
	Dept  string
	Age   int64
	Email string
	Tel   string
	Phone string
}
