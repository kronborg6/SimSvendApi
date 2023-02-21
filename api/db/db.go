package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// here we connet to the database useing a conntion string
func Init() *gorm.DB {
	// dsn := "root:Password@tcp(localhost:3306)/simsvend?charset=utf8mb4&parseTime=True&loc=Local" // Home Database
	dsn := "root:jxgcsTKUUO9OzYYN1Z56@tcp(containers-us-west-191.railway.app:6263)/railway?charset=utf8mb4&parseTime=True&loc=Local" // Home Database

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Setyup(db *gorm.DB) {
	db.AutoMigrate()
}
