package models

type Clubs struct {
	ID         int
	Name       string  `json:"name"`
	StreetName string  `json:"street_name"`
	Zipcode    int     `json:"ZipCode"`
	Zip        ZipCode `gorm:"foreignKey:Zipcode"`
}

type ClubCourts struct {
	ID     int
	Name   string `json:"name"`
	Double bool   `json:"single_court" gorm:"not null; default:true"`
	ClubId int    `json:"club_id"`
	Club   Clubs  `gorm:"foreignKey:ClubId"`
}
