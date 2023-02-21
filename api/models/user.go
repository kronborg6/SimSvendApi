package models

import "time"

type UserInfo struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" gorm:"unique not null"`
	Password    string `json:"password" gorm:"not null"`
	CreateAt    time.Time
	UserStatsID int64     `json:"user_stats_id"`
	UserStats   UserStats `gorm:"foreignKey:UserStatsID"`
	RoleId      int       `json:"role_id"`
	Role        Roles     `gorm:"foreignKey:RoleId"`
}

type UserStats struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	// UserInfoId int64    `json:"user_info_id"`
	// User       UserInfo `gorm:"foreignKey:UserInfoId"`
	Elo    int64 `json:"elo" gorm:"default:1000"`
	Points int64 `json:"points" gorm:"default:0"`
	Wins   int64 `json:"Wins" gorm:"default:0"`
	Losses int64 `json:"losses" gorm:"default:0"`
}

type Roles struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
