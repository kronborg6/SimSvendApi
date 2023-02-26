package repos

import (
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type CasualRepo struct {
	db *gorm.DB
}

func (repo *MatchRepo) NewCasualGame(game models.Matchs) (models.Matchs, error) {
	// var game models.Matchs
	if err := repo.db.Create(&game).Error; err != nil {
		return game, err
	}
	return game, nil
}
func (repo *MatchRepo) AddResualtCasualGame(game models.Matchs) (models.Matchs, error) {
	if err := repo.db.Model(&game).Where("id = ?", game.ID).Updates(&game).Error; err != nil {
		return game, nil
	}
	return game, nil

}
