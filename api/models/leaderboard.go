package models

type Leaderboards struct {
}

type Months struct {
	Id   int32  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
