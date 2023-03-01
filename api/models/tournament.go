package models

import "gorm.io/gorm"

type TournamentRepo struct {
	db *gorm.DB
}

func NewTournamentRepo(db *gorm.DB) *TournamentRepo {
	return &TournamentRepo{db}
}
