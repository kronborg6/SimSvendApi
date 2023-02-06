package models

type Clubs struct {
	Id int64 `json:"id" gorm:"primaryKey"`
}

type ClubCourts struct {
	Id int64 `json:"id" gorm:"primaryKey"`
}
