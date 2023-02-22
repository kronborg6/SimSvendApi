package models

import "time"

type UserInfo struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" gorm:"unique not null"`
	Password    string `json:"password" gorm:"not null"`
	CreateAt    time.Time
	UserStatsID int       `json:"user_stats_id"`
	UserStats   UserStats `gorm:"foreignKey:UserStatsID"`
	// FriendId    *int       `gorm:"default:null"`
	// FriendsList []UserInfo `gorm:"foreignKey:FriendsId"`
	RoleId int   `json:"role_id"`
	Role   Roles `gorm:"foreignKey:RoleId"`
}

type UserStats struct {
	Id int `json:"id" gorm:"primaryKey"`
	// UserInfoId int64    `json:"user_info_id"`
	// User       UserInfo `gorm:"foreignKey:UserInfoId"`
	Elo    int `json:"elo" gorm:"default:1000"`
	Points int `json:"points" gorm:"default:0"`
	Wins   int `json:"Wins" gorm:"default:0"`
	Losses int `json:"losses" gorm:"default:0"`
}

type Roles struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type Friends struct {
	Id     int      `json:"id" gorm:"primaryKey"`
	UserId int      `json:"user_id"`
	User   UserInfo `gorm:"foreignKey:UserId"`
}

type User struct {
	Id          int `json:"id" gorm:"primaryKey"`
	UserInfoId  int
	Userinfo    UserInfo `gorm:"foreignKey:UserInfoId"`
	UserStatsId int
	UserStats   UserStats `gorm:"foreignKey:UserStatsId"`
	// FriendList []Friends
	RoleId int
	Role   Roles `gorm:"foreignKey:RoleId"`
}
