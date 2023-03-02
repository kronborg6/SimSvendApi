package repos

import (
	"errors"

	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err := repo.db.Where("id = ?", id).Preload("Tour").Preload("Place").Preload("Players." + clause.Associations).Find(&tour)
	if err.Error != nil {
		return nil, errors.New("can't find tournamnet")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find tournamnet")
	}
	for i := range tour {
		for x := range tour[i].Players {
			tour[i].Players[x].Userinfo.Password = ""
			tour[i].Players[x].Userinfo.Email = ""
			tour[i].Players[x].Role.Name = ""
			tour[i].Players[x].Role.ID = 0
			tour[i].Players[x].RoleID = 0
			tour[i].Players[x].FriendList = nil
		}
	}
	return tour, nil
}

func NewTournamentRepo(db *gorm.DB) *TournamentRepo {
	return &TournamentRepo{db}
}
