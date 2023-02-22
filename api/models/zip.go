package models

type ZipCode struct {
	Id   int    `gorm:"primaryKey;autoIncrement:false"`
	Name string `json:"name"`
}
