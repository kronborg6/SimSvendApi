package models

import "time"

// type UserInfo struct {
// 	ID        int
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Email     string `json:"email" gorm:"unique; not null"`
// 	Password  string `json:"password" gorm:"not null"`
// 	CreateAt  time.Time
// 	// UserID    int `json:"userId"`
// 	// UserStatsID int       `json:"user_stats_id"`
// 	// UserStats   UserStats `gorm:"foreignKey:UserStatsID"`
// 	// FriendId    *int       `gorm:"default:null"`
// 	// FriendsList []UserInfo `gorm:"foreignKey:FriendsId"`
// 	// RoleId int   `json:"role_id"`
// 	// Role   Roles `gorm:"foreignKey:RoleId"`
// }

type UserInfo struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique; not null"`
	Password  string `json:"password" gorm:"not null"`
	CreateAt  time.Time
	UserID    int `json:"userId" gorm:"ForeignKey:ID"`
}

type UserStats struct {
	ID int
	// UserInfoId int64    `json:"user_info_id"`
	// User       UserInfo `gorm:"foreignKey:UserInfoId"`
	Elo    int `json:"elo" gorm:"default:1000"`
	Points int `json:"points" gorm:"default:0"`
	Wins   int `json:"Wins" gorm:"default:0"`
	Losses int `json:"losses" gorm:"default:0"`
	UserID int `json:"userId" gorm:"ForeignKey:ID"`
}

type Roles struct {
	ID   int
	Name string `json:"name" gorm:"not null"`
}

type Friends struct {
	ID int
	// UserInfoID int      `json:"UserInfoID"`
	// User       UserInfo `gorm:"foreignKey:UserInfoID"`
	UserInfoID int `json:"userId" gorm:"ForeignKey:ID"`
}

// type User struct {
// 	Id          int `json:"id" gorm:"primaryKey"`
// 	UserInfoId  int
// 	Userinfo    UserInfo `gorm:"foreignKey:UserInfoId"`
// 	UserStatsId int
// 	UserStats   UserStats `gorm:"foreignKey:UserStatsId"`
// 	FriendList  []Friends
// 	RoleId      int
// 	Role        Roles `gorm:"foreignKey:RoleId"`
// }

type User struct {
	ID int
	// UserInfoId  int
	Userinfo UserInfo `gorm:"NOT NULL" json:"userInfo"`
	// UserStatsId int
	UserStats UserStats
	// FriendList []Friends
	// RoleId      int
	// Role Roles
}
