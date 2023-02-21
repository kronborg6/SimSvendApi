package db

import (
	"fmt"
	"log"

	"github.com/kronborg6/SimSvendApi/api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func Init() *gorm.DB {
// 	dsn := "root:Password@tcp(localhost:3306)/simsvend?charset=utf8mb4&parseTime=True&loc=Local" // Home Database
// 	// dsn := "root:jxgcsTKUUO9OzYYN1Z56@tcp(containers-us-west-191.railway.app:6263)/railway?charset=utf8mb4&parseTime=True&loc=Local" // Home Database

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

func Init() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("MYSQLUSER"),
		config.Config("MYSQLPASSWORD"),
		config.Config("MYSQLHOST"),
		config.Config("MYSQLPORT"),
		config.Config("MYSQLDATABASE"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}

	return db

}

func Setyup(db *gorm.DB) {
	db.AutoMigrate()
}
