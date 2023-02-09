package models

type ZipCode struct {
	Id   int64  `gorm:"primaryKey;autoIncrement:false"`
	Name string `json:"name"`
}
