package repos

import (
	"gorm.io/gorm"
)

type LeaderboardRepo struct {
	db *gorm.DB
}

// func (repo *LeaderboardRepo) FindAllLeaderboards() ([]models.Leaderboards, error) {}
// func (repo *LeaderboardRepo) FindLeaderboards() ([]models.Leaderboards, error)    {}
