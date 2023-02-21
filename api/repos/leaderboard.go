package repos

import (
	"errors"
	"fmt"

	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type LeaderboardRepo struct {
	db *gorm.DB
}

func (repo *LeaderboardRepo) FindAllLeaderboards(mounthId int) ([]models.Leaderboards, error) {
	var leaderboard []models.Leaderboards
	err := repo.db.Debug().Where("month_id = ?", mounthId).Preload("Player").Preload("Month").Find(&leaderboard)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't a Leaderboard")
	}

	leaderboard[0].Player.Password = ""
	fmt.Println(len(leaderboard))
	return leaderboard, nil
}
func (repo *LeaderboardRepo) FindLeaderboards() ([]models.Leaderboards, error) {
	return nil, nil

}

func NewLeaderbaordRepo(db *gorm.DB) *LeaderboardRepo {
	return &LeaderboardRepo{db}
}
