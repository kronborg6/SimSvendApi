package models

import (
	"time"

	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Clubs{},
		&Months{},
		&ClubCourts{},
		&Leaderboards{},
		&Matchs{},
		&Results{},
		&Roles{},
		&ZipCode{},
		&UserInfo{},
		&UserStats{},
	)
	db.AutoMigrate(
		&Clubs{},
		&Months{},
		&ClubCourts{},
		&Leaderboards{},
		&Matchs{},
		&Results{},
		&Roles{},
		&ZipCode{},
		&UserInfo{},
		&UserStats{},
	)

	role := []Roles{
		{
			Name: "Bruger",
		},
		{
			Name: "Admin",
		},
	}
	userinfo := []UserInfo{
		{
			FirstName: "Mikkel",
			LastName:  "Kronborg",
			Email:     "mkronborg7@gmail.com",
			Password:  "Test",
			RoleId:    2,
			CreateAt:  time.Now(),
		},
		{
			FirstName: "August",
			LastName:  "Schnell",
			Email:     "augustschnellpedersen@gmail.com",
			Password:  "Test",
			RoleId:    1,
			CreateAt:  time.Now(),
		},
	}
	db.Create(role)
	db.Create(userinfo)
}
