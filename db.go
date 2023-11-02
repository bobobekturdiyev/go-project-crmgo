package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

const (
	DBNAME     = "go_db"
	DBUSER     = "root"
	DBPASSWORD = ""
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASSWORD, DBNAME)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Auto migrate will create the table and missing columns, it won't change existing columns' type or delete unused columns
	db.AutoMigrate(&User{})
}
