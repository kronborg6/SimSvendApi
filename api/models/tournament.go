package models

import "time"

type Tournament struct {
	ID      int
	Name    string    `json:"name"`
	HowMany int       `json:"how_many"`
	PlaceID int       `json:"place_id"`
	Place   Club      `gorm:"foreignkey:PlaceID" `
	Date    time.Time `json:"date"`
	Elo     int       `json:"elo"`
	Gender  string    `json:"gender"`
	Players []User    `gorm:"many2many:tournament_players;"`
}

type TournamentInfo struct {
	ID int
}
