package models

import "time"

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
	ID     int
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
	ID        int
	UserID    int  `gorm:"not null"`
	FriendID  int  `json:"friend_id" gorm:"NOT NULL"`
	Friend    User `gorm:"foreignkey:FriendID" sql:"DEFAULT:null"`
	IsFriends bool `json:"is_friends" sql:"DEFAULT:false"`
	Sender    bool `json:"-" sql:"DEFUALT:false"`
}

type User struct {
	ID         int
	Userinfo   UserInfo `gorm:"NOT NULL" json:"userInfo"`
	UserStats  UserStats
	RoleID     int
	Role       Roles     `gorm:"foreignkey:RoleID" `
	FriendList []Friends `gorm:"foreignkey:UserID" sql:"DEFAULT:null"`
}
