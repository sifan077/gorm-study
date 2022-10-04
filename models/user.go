package models

import "gorm.io/gorm"

type User struct {
	Id     int `gorm:"column:id;primary_key;AUTO_INCREMENT" `
	Name   string
	Gender string
	Hobby  string
	Email  string
}

type Demo struct {
	Id      int `gorm:"column:id;primary_key;AUTO_INCREMENT" `
	Name    string
	Deleted gorm.DeletedAt
}
