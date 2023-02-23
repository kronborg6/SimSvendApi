package models

type Leaderboards struct {
	ID       int
	MonthID  int    `json:"month_id"`
	Month    Months `gorm:"foreignKey:MonthID"`
	Year     int32
	PlayerID int
	Player   UserInfo `gorm:"foreignKey:PlayerID"`
	Points   int      `gorm:"default:null"`
}

type Months struct {
	ID   int
	Name string `json:"name"`
}
