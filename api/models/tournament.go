package models

import "time"

type Tournament struct {
	ID      int
	Name    string    `json:"name" gorm:"NOT NULL"`
	HowMany int       `json:"how_many" gorm:"NOT NULL"`
	PlaceID int       `json:"place_id" gorm:"NOT NULL"`
	Place   Club      `gorm:"foreignkey:PlaceID" `
	Date    time.Time `json:"date" gorm:"NOT NULL"`
	Elo     int       `json:"elo" gorm:"NOT NULL"`
	Gender  string    `json:"gender" gorm:"NOT NULL"`
	Tour    TournamentInfo
	Players []User `gorm:"many2many:tournament_players;"`
}

type TournamentInfo struct {
	ID           int
	PricePool    int    `json:"price_pool" gorm:"NOT NULL"`
	Dec          string `json:"dec" gorm:"NOT NULL"`
	TournamentID int    `json:"tournament_id" gorm:"NOT NULL"`
}

type JoinTourModel struct {
	UserID int `json:"user_id"`
	TourID int `json:"tour_id"`
}
