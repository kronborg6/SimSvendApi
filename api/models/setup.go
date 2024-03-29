package models

import (
	"fmt"
	"time"

	"github.com/kronborg6/SimSvendApi/api/middleware"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&ClubCourts{},
		&Club{},
		&Months{},
		&Results{},
		&Match{},
		&Roles{},
		&ZipCode{},
		&User{},
		&UserStats{},
		&UserInfo{},
		&Friends{},
		&Leaderboards{},
		&TournamentInfo{},
		&Tournament{},
	)
	db.AutoMigrate(
		&ClubCourts{},
		&Club{},
		&Months{},
		&Results{},
		&Match{},
		&Roles{},
		&ZipCode{},
		&User{},
		&UserStats{},
		&UserInfo{},
		&Friends{},
		&Leaderboards{},
		&TournamentInfo{},
		&Tournament{},
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
	// 		FirstName: "kristian",
	// 		LastName:  "Kronborg",
	// 		Email:     "kronborg@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    2,
	// 		// UserStatsID: 1,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 2}, {Id: 3}},
	// 	},
	// 	{
	// 		FirstName: "August",
	// 		LastName:  "Jensen",
	// 		Email:     "AugJensen@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    1,
	// 		// UserStatsID: 2,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 1}},
	// 	},
	// 	{
	// 		FirstName: "Ole",
	// 		LastName:  "Olesen",
	// 		Email:     "oleole@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		// RoleId:    1,
	// 		// UserStatsID: 3,
	// 		CreateAt: time.Now(),
	// 		// FriendsList: []UserInfo{{Id: 1}, {Id: 4}},
	// 	},
	// 	{
	// 		FirstName: "Morten",
	// 		LastName:  "Iorden",
	// 		Email:     "morteeen@gmail.com",
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

	// user := []User{
	// 	{

	// 		Userinfo: UserInfo{
	// 			FirstName: "kristian",
	// 			LastName:  "Kronborg",
	// 			Email:     "Kronborg@gmail.com",
	// 			Password:  middleware.HashPassword("Test"),
	// 			CreateAt:  time.Now(),
	// 		},
	// 	},

	// {

	// 	Userinfo: UserInfo{
	// 		FirstName: "Ole",
	// 		LastName:  "Ost",
	// 		Email:     "Osten@gmail.com",
	// 		Password:  middleware.HashPassword("Test"),
	// 		CreateAt:  time.Now(),
	// 	},
	// },
	// {
	// 	Userinfo: UserInfo{
	// 		FirstName: "Jens",
	// 		LastName:  "Jensen",
	// 		Password:  middleware.HashPassword("Test"),
	// 		Email:     "j.jensen@gmail.com",
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
			FirstName: "Morten",
			LastName:  "Hansen",
			Password:  middleware.HashPassword("Test"),
			Email:     "t.hansen@gmail.com",
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
	// }

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
			MonthID:  1,
			Year:     2023,
			PlayerID: 1,
			Points:   6996,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 2,
			Points:   1,
		},
		{
			MonthID:  1,
			Year:     2023,
			PlayerID: 3,
			Points:   221,
		},
		/* {
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
		}, */
	}
	zip := []ZipCode{
		{
			ID:   6996,
			Name: "Fucking",
		},
		{
			ID:   5000,
			Name: "Broby",
		},
		{
			ID:   4200,
			Name: "Snoop",
		},
	}
	user2 := []User{
		{
			Userinfo: UserInfo{
				FirstName: "kristian",
				LastName:  "Kronborg",
				Password:  middleware.HashPassword("Test"),
				Email:     "k.kronborg@gmail.com",
				CreateAt:  time.Now(),
			},
			RoleID:    2,
			UserStats: UserStats{Elo: 99999, Points: 1241241, Wins: 999, Losses: 0},
		},
		{

			Userinfo: UserInfo{
				FirstName: "Tina",
				LastName:  "Tomsen",
				Email:     "t.tomsen@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
			RoleID:    1,
			UserStats: UserStats{Elo: 150, Points: 250, Wins: 0, Losses: 0},
		},
		{

			Userinfo: UserInfo{
				FirstName: "Allan",
				LastName:  "hansen",
				Email:     "hansenerbaresej@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
			RoleID:    1,
			UserStats: UserStats{Elo: 423, Points: 5555, Wins: 10, Losses: 0},
		},
		{

			Userinfo: UserInfo{
				FirstName: "thomans",
				LastName:  "hansen",
				Email:     "t.hansen@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
			RoleID:    1,
			UserStats: UserStats{Elo: 100, Points: 1, Wins: 0, Losses: 100},
		},
		{

			Userinfo: UserInfo{
				FirstName: "Oliver",
				LastName:  "Mikkelsen",
				Email:     "Oliver69@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
			RoleID:    1,
			UserStats: UserStats{Elo: 100, Points: 1, Wins: 0, Losses: 100},
		},
		{

			Userinfo: UserInfo{
				FirstName: "Morten",
				LastName:  "Iorden",
				Email:     "deterbareiorden@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
			RoleID:    1,
			UserStats: UserStats{Elo: 100, Points: 1, Wins: 0, Losses: 100},
		},
		{

			Userinfo: UserInfo{
				FirstName: "Jakob",
				LastName:  "Jakobson",
				Email:     "JakobJohansen@gmail.com",
				Password:  middleware.HashPassword("Test"),
				CreateAt:  time.Now(),
			},
			RoleID:    1,
			UserStats: UserStats{Elo: 100, Points: 1, Wins: 0, Losses: 100},
		},
	}
	place := []Club{
		{
			Name: "Padelboxen",
			// Zipcode:    5000,
			StreetName: "Tolderlundsvej 92",
			Courts:     []ClubCourts{{Name: "CC", Double: true}, {Name: "D1", Double: true}},
		},
		{
			Name: "PadelPadel",
			// Zipcode:    5220,
			StreetName: "C. F. Tietgens Blvd. 24K",
			Courts:     []ClubCourts{{Name: "S1", Double: false}, {Name: "S2", Double: false}},
		},
	}
	// courts := []ClubCourts{
	// 	{
	// 		Name:   "CC",
	// 		Double: true,
	// 		// ClubId: 1,
	// 	},
	// 	{
	// 		Name:   "D1",
	// 		Double: true,
	// 		// ClubId: 1,
	// 	},
	// 	{
	// 		Name:   "D2",
	// 		Double: true,
	// 		// ClubId: 1,
	// 	},
	// 	{
	// 		Name:   "D3",
	// 		Double: true,
	// 		// ClubId: 1,
	// 	},
	// 	{
	// 		Name:   "S1",
	// 		Double: false,
	// 		// ClubId: 1,
	// 	},
	// 	{
	// 		Name:   "S2",
	// 		Double: false,
	// 		// ClubId: 2,
	// 	},
	// 	{
	// 		Name:   "S1",
	// 		Double: false,
	// 		// ClubId: 2,
	// 	},
	// }

	// Tjek om der er nogen fejl under oprettelse
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// 	// Håndter fejl
	// }
	// db.Create(&user)
	// db.Create(&user2)
	matchs := []Match{
		{
			PlayTime:     time.Now(),
			PlaceId:      1,
			CourtId:      1,
			TeamAPlayerA: 1,
			TeamAPlayerB: 2,
			TeamBPlayerA: 3,
			TeamBPlayerB: 4,
			// Don:          false,
			// Result:       Results{},
			Result: &Results{},
		},
		{
			PlayTime:     time.Now(),
			PlaceId:      1,
			CourtId:      1,
			TeamAPlayerA: 3,
			TeamAPlayerB: 1,
			TeamBPlayerA: 2,
			TeamBPlayerB: 4,
			Don:          true,
			// Result:       Results{},
			Result: &Results{AOne: 6, BOne: 0, ATwo: 6, BTwo: 2},
		},

		{
			PlayTime:     time.Now(),
			PlaceId:      1,
			CourtId:      1,
			TeamAPlayerA: 2,
			TeamAPlayerB: 3,
			TeamBPlayerA: 4,
			TeamBPlayerB: 1,
			Don:          true,
			// Result:       Results{},
			Result: &Results{AOne: 6, BOne: 4, ATwo: 3, BTwo: 6, AThree: 4, BThree: 6},
		},
		// {
		// 	PlayTime:     time.Now(),
		// 	PlaceId:      1,
		// 	CourtId:      1,
		// 	TeamAPlayerA: 1,
		// 	TeamAPlayerB: 2,
		// 	TeamBPlayerA: 3,
		// 	TeamBPlayerB: 4,
		// 	Don:          false,
		// 	Result:       Results{},
		// },
		// {
		// 	PlayTime:     time.Now(),
		// 	PlaceId:      1,
		// 	CourtId:      1,
		// 	TeamAPlayerA: 1,
		// 	TeamAPlayerB: 2,
		// 	TeamBPlayerA: 3,
		// 	TeamBPlayerB: 4,
		// 	Don:          false,
		// 	Result:       Results{},
		// },
		// {
		// 	PlayTime:     time.Now(),
		// 	PlaceId:      1,
		// 	CourtId:      1,
		// 	TeamAPlayerA: 1,
		// 	TeamAPlayerB: 2,
		// 	TeamBPlayerA: 3,
		// 	TeamBPlayerB: 4,
		// 	Don:          false,
		// 	Result:       Results{},
		// },
	}
	tour := []Tournament{
		{
			Name:    "King of Padel",
			HowMany: 3,
			PlaceID: 1,
			Date:    time.Now(),
			Gender:  "Men",
			Players: []User{{ID: 1}, {ID: 3}},
			Tour:    TournamentInfo{PricePool: 1000, Dec: "Vis i tro i er de best padel par så find ud af det her"},
		},
		{
			Name:    "Queen of Padel",
			HowMany: 2,
			PlaceID: 1,
			Date:    time.Now(),
			Gender:  "Women",
			Players: []User{{ID: 2}, {ID: 4}},
			Tour:    TournamentInfo{PricePool: 1000, Dec: "Vis i tro i er de best padel par så find ud af det her"},
		},
	}
	friends := []Friends{
		{
			UserID:    1,
			FriendID:  2,
			IsFriends: true,
		},
		{
			UserID:    2,
			FriendID:  1,
			IsFriends: true,
		},
		{
			UserID:    1,
			FriendID:  6,
			IsFriends: true,
		},

		{
			UserID:    6,
			FriendID:  1,
			IsFriends: true,
		},

		{
			UserID:    1,
			FriendID:  7,
			IsFriends: true,
		},

		{
			UserID:    7,
			FriendID:  1,
			IsFriends: true,
		},
		{
			UserID:    7,
			FriendID:  3,
			IsFriends: true,
		},
		{
			UserID:    3,
			FriendID:  7,
			IsFriends: true,
		},
		{
			UserID:    3,
			FriendID:  4,
			IsFriends: true,
		},
		{
			UserID:    3,
			FriendID:  5,
			IsFriends: true,
		},
		{
			UserID:    5,
			FriendID:  3,
			IsFriends: true,
		},
		{
			UserID:    4,
			FriendID:  3,
			IsFriends: true,
		},
		{
			UserID:    6,
			FriendID:  3,
			IsFriends: true,
		},
		{
			UserID:    3,
			FriendID:  6,
			IsFriends: true,
		},

		// {
		// 	UserID:    1,
		// 	FriendID:  3,
		// 	IsFriends: false,
		// },
		// {
		// 	UserID:    3,
		// 	FriendID:  1,
		// 	IsFriends: false,
		// },
	}

	db.Create(&role)

	for i := range user2 {
		if err := db.Create(&user2[i]).Error; err != nil {
			fmt.Println("fuck nej")
		}
	}
	for i := range place {
		if err := db.Create(&place[i]).Error; err != nil {
			fmt.Println("fuck nej")
		}
	}
	// for i := range matchs {
	// 	if err := db.Create(&matchs[i]).Error; err != nil {
	// 		fmt.Println("hej med dig")
	// 	}
	// }
	// db.Create(userStats)
	// db.Create(userinfo)
	// db.Create(friends)
	db.Create(&months)
	db.Create(&zip)
	// db.Create(&courts)
	// db.Create(&place)
	db.Create(&matchs)
	db.Create(&leaderboards)
	db.Create(&tour)
	db.Create(&friends)

	/*
		 	err := db.SetupJoinTable(&Tournament{}, "test", &User{})

			if err != nil {
				fmt.Println(err)
			}
	*/
}
