package models

type ZipCode struct {
	ID   int    `gorm:"primaryKey;autoIncrement:false"`
	Name string `json:"name"`
}
