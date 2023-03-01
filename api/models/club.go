package models

type Club struct {
	ID         int
	Name       string `json:"name"`
	StreetName string `json:"street_name"`
	// Zipcode    int    `json:"ZipCode"`
	// Zip        ZipCode `gorm:"foreignKey:Zipcode"`
	Courts []ClubCourts
}

type ClubCourts struct {
	ID     int
	Name   string `json:"name"`
	Double bool   `json:"single_court" gorm:"not null; default:true"`
	ClubID int    `gorm:"ForeignKey:ID"`
	// ClubId int    `json:"club_id"`
	// Club   Clubs  `gorm:"foreignKey:ClubId"`
}
