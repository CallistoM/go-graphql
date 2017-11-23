package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       int    `json:”id”`
	Email    string `json:email`
	Surname  string `json:”surname”`
	Lastname string `json:”lastname”`
}
