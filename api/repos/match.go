package repos

import (
	"errors"

	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type MatchRepo struct {
	db *gorm.DB
}

func (repo *MatchRepo) FindMyGame(gameID int) ([]models.Matchs, error) {
	var game []models.Matchs

	err := repo.db.Where("id = ?", gameID).Find(&game)
	if err.Error != nil {
		return nil, errors.New("can't find game")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find game")
	}

	return game, nil
}
func (repo *MatchRepo) FindMyGames(userID int) ([]models.Matchs, error) {
	var games []models.Matchs

	err := repo.db.Where("team_a_player_a = ?", userID).Or("team_a_player_b = ?", userID).Or("team_b_player_a = ?", userID).Or("team_b_player_b = ?", userID).Find(&games)
	if err.Error != nil {
		return nil, errors.New("can't find games")
	}

	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find games")
	}
	return games, nil
}
func (repo *MatchRepo) GameHistory() ([]models.Matchs, error)
func (repo *MatchRepo) NewCasualGame() (models.Matchs, error)

func NewMatchRepo(db *gorm.DB) *MatchRepo {
	return &MatchRepo{db}
}
