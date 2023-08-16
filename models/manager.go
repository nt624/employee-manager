package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model

	ID   int64
	Name string
	Dept string
	Age  int64
}
