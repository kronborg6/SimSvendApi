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

	userStats := []UserStats{
		{
			UserInfoId: 1,
			Elo:        99999,
			Points:     142347,
			Wins:       1000,
			Losses:     0,
		},
		{
			UserInfoId: 2,
			Elo:        -4532,
			Points:     12,
			Wins:       2,
			Losses:     2212,
		},
	}
	months := []Months{
		{
			Name: "Januar",
		},
		{
			Name: "Februar",
		},
		{
			Name: "Marts",
		},
		{
			Name: "April",
		},
		{
			Name: "Maj",
		},
		{
			Name: "Juni",
		},
		{
			Name: "Juli",
		},
		{
			Name: "August",
		},
		{
			Name: "September",
		},
		{
			Name: "Oktober",
		},
		{
			Name: "November",
		},
		{
			Name: "December",
		},
	}
	leaderboards := []Leaderboards{
		{

		}
	}
	zip := []ZipCode{
		{
			Id:   5690,
			Name: "Tommerup",
		},
		{
			Id:   5000,
			Name: "Odense",
		},
		{
			Id:   5220,
			Name: "Odense",
		},
	}
	place := []Clubs{
		{
			Name:       "Padelboxen",
			Zipcode:    5000,
			StreetName: "Tolderlundsvej 92",
		},
		{
			Name:       "PadelPadel",
			Zipcode:    5220,
			StreetName: "C. F. Tietgens Blvd. 24K",
		},
	}
	courts := []ClubCourts{
		{
			Name:   "CC",
			Double: true,
			ClubId: 1,
		},
		{
			Name:   "D1",
			Double: true,
			ClubId: 1,
		},
		{
			Name:   "D2",
			Double: true,
			ClubId: 1,
		},
		{
			Name:   "D3",
			Double: true,
			ClubId: 1,
		},
		{
			Name:   "S1",
			Double: false,
			ClubId: 1,
		},
		{
			Name:   "S2",
			Double: false,
			ClubId: 2,
		},
		{
			Name:   "S1",
			Double: false,
			ClubId: 2,
		},
	}
	db.Create(role)
	db.Create(userinfo)
	db.Create(userStats)
	db.Create(months)
	db.Create(zip)
	db.Create(place)
	db.Create(courts)
}
