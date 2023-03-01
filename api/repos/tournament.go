package repos

import (
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type TournamentRepo struct {
	db *gorm.DB
}

func (repo *TournamentRepo) FindAllTour() ([]models.Tournament, error) {
	var tour []models.Tournament

	if err := repo.db.Debug().Preload("Place").Find(&tour).Error; err != nil {
		return nil, err
	}
	return tour, nil
}

func (repo *TournamentRepo) FindTour() ([]models.Tournament, error) {
	return nil, nil
}

func NewTournamentRepo(db *gorm.DB) *TournamentRepo {
	return &TournamentRepo{db}
}
