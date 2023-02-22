package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "root:Password@tcp(localhost:3306)/simsvend?charset=utf8mb4&parseTime=True&loc=Local" // Home Database
	// dsn := "root:jxgcsTKUUO9OzYYN1Z56@tcp(containers-us-west-191.railway.app:6263)/railway?charset=utf8mb4&parseTime=True&loc=Local" // Home Database

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}

// func Init() *gorm.DB {
// 	fmt.Println("dfgdfgdgf")
// 	fmt.Println(os.Getenv("hej"))
// 	dsn := fmt.Sprintf(
// 		"%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
// 		os.Getenv("MYSQLUSER"),
// 		os.Getenv("MYSQLPASSWORD"),
// 		os.Getenv("MYSQLHOST"),
// 		os.Getenv("MYSQLPORT"),
// 		os.Getenv("MYSQLDATABASE"),
// 	)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err)
// 		panic("Failed to connect to database")
// 	}

// 	return db

// }

func Setyup(db *gorm.DB) {
	db.AutoMigrate()
}
