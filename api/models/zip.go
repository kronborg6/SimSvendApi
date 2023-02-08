package models

type ZipCode struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
