package models

import (
	"time"

	"github.com/kronborg6/SimSvendApi/api/middleware"
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
		&User{},
		&UserStats{},
		&UserInfo{},
		&Friends{},
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
		&User{},
		&UserStats{},
		&UserInfo{},
		&Friends{},
	)

	role := []Roles{
		{
			Name: "Bruger",
		},
		{
			Name: "Admin",
		},
	}
	// userinfo := []UserInfo{
	// 	{
	// 		FirstName: "Mikkel",
	// 		LastName:  "Kronborg",
	// 		Email:     "mkronborg7@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    2,
	// 		// UserStatsID: 1,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 2}, {Id: 3}},
	// 	},
	// 	{
	// 		FirstName: "August",
	// 		LastName:  "Schnell",
	// 		Email:     "augustschnellpedersen@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    1,
	// 		// UserStatsID: 2,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 1}},
	// 	},
	// 	{
	// 		FirstName: "Tina",
	// 		LastName:  "Kronborg",
	// 		Email:     "t.kronborg6@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    1,
	// 		// UserStatsID: 3,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 1}, {Id: 4}},
	// 	},
	// 	{
	// 		FirstName: "Oliver",
	// 		LastName:  "Mathiesen",
	// 		Email:     "mathiesenoliver@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    1,
	// 		// UserStatsID: 4,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 2}, {Id: 3}},
	// 	},
	// }
	// friends := []Friends{
	// 	{
	// 		UserId: 1,
	// 	},
	// 	{
	// 		UserId: 2,
	// 	},
	// 	{
	// 		UserId: 3,
	// 	},
	// 	{
	// 		UserId: 4,
	// 	},
	// }

	// userStats := []UserStats{
	// 	{
	// 		Elo:    99999,
	// 		Points: 142347,
	// 		Wins:   1000,
	// 		Losses: 0,
	// 	},
	// 	{
	// 		Elo:    -4532,
	// 		Points: 12,
	// 		Wins:   2,
	// 		Losses: 2212,
	// 	},
	// 	{
	// 		Elo:    1206,
	// 		Points: 13021,
	// 		Wins:   764,
	// 		Losses: 453,
	// 	},
	// 	{
	// 		Elo:    5933,
	// 		Points: 9542,
	// 		Wins:   765,
	// 		Losses: 467,
	// 	},
	// }

	user := []User{
		{

			Userinfo: UserInfo{
				FirstName: "Mikkel",
				LastName:  "Kronborg",
				Email:     "mkronborg7@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
		},
		// {

		// 	Userinfo: UserInfo{
		// 		FirstName: "Tina",
		// 		LastName:  "Kronborg",
		// 		Email:     "t.kronborg6@gmail.com",
		// 		Password:  middleware.HashPassword("Test"),
		// 		CreateAt:  time.Now(),
		// 	},
		// },
		// {
		// 	Userinfo: UserInfo{
		// 		FirstName: "Mikkel",
		// 		LastName:  "Kronborg",
		// 		Password:  middleware.HashPassword("Test"),
		// 		Email:     "mkronborg7@gmail.com",
		// 		CreateAt:  time.Now(),
		// 	},
		// 	// UserStats: UserStats{},
		// 	// Role:      Roles{Name: "Admin"},
		// 	// UserInfoId:  1,
		// 	// UserStatsId: 1,
		// 	// RoleId:      1,
		// 	// FriendList: []Friends{{UserID: 2}},
		// },
		/* 		{
			Userinfo: UserInfo{
				FirstName: "Tina",
				LastName:  "Kronborg",
				Password:  middleware.HashPassword("Test"),
				Email:     "t.kronborg6@gmail.com",
				CreateAt:  time.Now(),
			},
			// 	// UserStats: UserStats{},
			// 	// Role:      Roles{Name: "Bruger"},
			// 	// UserInfoId:  1,
			// 	// UserStatsId: 1,
			// 	// RoleId:      1,
			// 	// FriendList: []Friends{{UserID: 1}},
		}, */
		// {
		// 	UserInfoId:  2,
		// 	UserStatsId: 2,
		// 	RoleId:      2,
		// 	FriendList:  []Friends{{UserId: 1}, {UserId: 4}},
		// },
		// {
		// 	UserInfoId:  3,
		// 	UserStatsId: 3,
		// 	RoleId:      2,
		// 	FriendList:  []Friends{{UserId: 1}, {UserId: 4}},
		// },
		// {
		// 	UserInfoId:  4,
		// 	UserStatsId: 4,
		// 	RoleId:      2,
		// 	FriendList:  []Friends{{UserId: 3}, {UserId: 2}},
		// },
		// {
		// 	UserInfoId:  1,
		// 	UserStatsId: 1,
		// 	// Userinfo:  userinfo[0],
		// 	// UserStats: userStats[0],
		// 	// Role:      role[0],
		// 	// FriendList: []Friends{{Id: 2}},
		// },
		// {
		// 	// Userinfo:  userinfo[1],
		// 	// UserStats: userStats[1],
		// 	// Role:      role[1],
		// 	UserInfoId:  2,
		// 	UserStatsId: 2,
		// 	// FriendList: []Friends{{Id: 1}},
		// },
		// {
		// 	// Userinfo:  userinfo[2],
		// 	// UserStats: userStats[2],
		// 	// Role:      role[2],
		// 	UserInfoId:  3,
		// 	UserStatsId: 3,
		// 	// FriendList: []Friends{{Id: 4}},
		// },
		// {
		// 	// Userinfo:  userinfo[3],
		// 	// UserStats: userStats[3],
		// 	// Role:      role[3],
		// 	UserInfoId:  3,
		// 	UserStatsId: 3,
		// 	// FriendList: []Friends{{Id: 3}},
		// },
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
	/* 	leaderboards := []Leaderboards{
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 1,
			Points:   6996,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 2,
			Points:   4223,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 3,
			Points:   221,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 4,
			Points:   420,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 1,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 2,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 3,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 4,
		},
	} */
	zip := []ZipCode{
		{
			ID:   5690,
			Name: "Tommerup",
		},
		{
			ID:   5000,
			Name: "Odense",
		},
		{
			ID:   5220,
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

	// Tjek om der er nogen fejl under oprettelse
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// 	// Håndter fejl
	// }
	db.Create(&user)
	db.Create(&role)
	// db.Create(userStats)
	// db.Create(userinfo)
	// db.Create(friends)
	db.Create(&months)
	db.Create(&zip)
	db.Create(&place)
	db.Create(&courts)
	// db.Create(&leaderboards)
}
