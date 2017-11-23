package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InsertUser(user *User) error {
	var id int
	err := db.Create(&user)

	fmt.Println(err)

	// err := db.QueryRow(`
	// 	INSERT INTO users(email)
	// 	VALUES ($1)
	// 	RETURNING id
	// `, user.Email).Scan(&id)
	user.ID = id
	return nil
}

func GetUserByID(id int) (*User, error) {
	var email string

	user := User{}

	err := db.Where("id = ?", id).First(&user)

	fmt.Println(err)

	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func GetUsers() ([]*User, error) {

	var allUsers []*User

	db.Find(&allUsers)

	return &allUsers, nil

}

func RemoveUserByID(id int) error {
	err := db.Where("id = ?", id).Delete(&User{})
	return err.Error
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
