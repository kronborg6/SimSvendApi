package repos

import (
	"errors"

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

func (repo *TournamentRepo) FindTour(id int) ([]models.Tournament, error) {
	var tour []models.Tournament
	err := repo.db.Where("id = ?", id).Preload("").Find(&tour)
	if err.Error != nil {
		return nil, errors.New("can't find tournamnet")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find tournamnet")
	}
	return tour, nil
}

func NewTournamentRepo(db *gorm.DB) *TournamentRepo {
	return &TournamentRepo{db}
}
