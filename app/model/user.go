package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Hash string
	Username string
	Password string
	Mobile string
}

