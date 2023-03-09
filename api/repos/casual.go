package repos

import (
	"errors"
	"time"

	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type CasualRepo struct {
	db *gorm.DB
}

func (repo *CasualRepo) NewCasualGame(game models.Match) (models.Match, error) {
	game.Result = &models.Results{}
	game.PlayTime = time.Now()
	if err := repo.db.Create(&game).Error; err != nil {
		return game, err
	}
	return game, nil
}
func (repo *CasualRepo) AddResualtCasualGame(Score models.Results) (models.Results, error) {
	var game models.Results
	var match models.Match
	if err := repo.db.Debug().Where("id = ?", Score.MatchID).Find(&match).Error; err != nil {
		return game, err
	}
	if match.Don {
		return game, errors.New("game is allready don")
	}
	if err := repo.db.Debug().Model(&Score).Where("match_id = ?", Score.MatchID).Updates(&Score).Error; err != nil {
		return game, err
	}
	match.Don = true
	if err := repo.db.Debug().Model(&match).Where("id = ?", Score.MatchID).Updates(&match).Error; err != nil {
		return game, err
	}
	return game, nil
}

func NewCasualRepo(db *gorm.DB) *CasualRepo {
	return &CasualRepo{db}
}
