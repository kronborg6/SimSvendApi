package models

import "gorm.io/gorm"

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
}
