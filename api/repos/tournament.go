package repos

import (
	"errors"
	"time"

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

func (repo *TournamentRepo) JoinTour(user models.JoinTourModel) (models.Tournament, error) {
	var tour models.Tournament
	var player []models.User
	if err := repo.db.Preload("Players").First(&tour, user.TourID).Error; err != nil {
		return tour, err
	}
	if err := repo.db.First(&player, user.UserID).Error; err != nil {
		return tour, err
	}
	if len(tour.Players) >= tour.HowMany {
		return tour, errors.New("max number off players")
	}
	for i := range tour.Players {
		if tour.Players[i].ID == user.UserID {
			return tour, errors.New("user is allready on the tour")
		}
	}
	tour.Players = append(tour.Players, player...)
	if err := repo.db.Debug().Save(&tour).Error; err != nil {
		return tour, err
	}

	return tour, nil
}
func (repo *TournamentRepo) NewTour(model models.Tournament) (models.Tournament, error) {
	var tour models.Tournament
	model.Date = time.Now()
	if err := repo.db.Debug().Where("id = ?", model.ID).Create(&model).Error; err != nil {
		return tour, err
	}
	return tour, nil
}
func (repo *TournamentRepo) UpdateTourInfo(info models.TournamentInfo) error {

	if err := repo.db.Debug().Model(&info).Where("tournament_id = ?", info.TournamentID).Updates(&info).Error; err != nil {
		return err
	}
	return nil
}
func (repo *TournamentRepo) UpdateTour(tour models.Tournament) error {
	// if err := repo.db.Debug().Model(&tour).Association("Tour").Append(&tour); err != nil {
	// 	return err
	// }
	// if err := repo.db.Debug().Model(&tour).Where("id = ?", tour.ID).Updates(&tour).Error; err != nil {
	// 	return err
	// }
	// return nil
	var t models.TournamentInfo
	var tourinfo = tour.Tour
	tour.Tour = t
	if err := repo.db.Debug().Model(&tour).Updates(&tour).Error; err != nil {
		return err
	}

	// update the TournamentInfo model
	if err := repo.db.Debug().Model(&t).Where("tournament_id = ?", tourinfo.TournamentID).Updates(&tourinfo).Error; err != nil {
		return err
	}
	// if err := repo.db.Debug().Model(&tour).Association("Tour").Replace(&tourinfo); err != nil {
	// 	return err
	// }

	return nil
}

// func (repo *TournamentRepo) UpdateTour(tour models.Tournament) error {
// 	if err := repo.db.Debug().Model(&tour).Updates(&tour).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func NewTournamentRepo(db *gorm.DB) *TournamentRepo {
	return &TournamentRepo{db}
}
