package models

type Leaderboards struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	MonthID  int    `json:"month_id"`
	Month    Months `gorm:"foreignKey:MonthID"`
	Year     int32
	PlayerID int
	Player   UserInfo `gorm:"foreignKey:PlayerID"`
	Points   int      `gorm:"default:null"`
}

type Months struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
