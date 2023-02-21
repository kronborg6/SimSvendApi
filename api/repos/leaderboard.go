package repos

import (
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type LeaderboardRepo struct {
	db *gorm.DB
}

func (repo *LeaderboardRepo) FindAllLeaderboards() ([]models.Leaderboards, error) {
	var leaderboard []models.Leaderboards
	var err error
	if err = repo.db.Debug().Preload("Player").Preload("Month").Find(&leaderboard).Error; err != nil {
		return leaderboard, err
	}
	leaderboard[0].Player.Password = ""
	return leaderboard, nil
}
func (repo *LeaderboardRepo) FindLeaderboards() ([]models.Leaderboards, error) {
	return nil, nil

}

func NewLeaderbaordRepo(db *gorm.DB) *LeaderboardRepo {
	return &LeaderboardRepo{db}
}
