package repos

import (
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type UserStatsRepo struct {
	db *gorm.DB
}

func (repo *UserStatsRepo) FindAllPlayerStats() ([]models.UserStats, error) {
	var userStats []models.UserStats

	if err := repo.db.Debug().Find(&userStats).Error; err != nil {
		return userStats, err
	}
	return userStats, nil
}

func (repo *UserStatsRepo) FindPlayerStats(email string) ([]models.UserStats, error) {
	var userStats []models.UserStats

	if err := repo.db.Debug().Where("email = ?", email).Find(userStats).Error; err != nil {
		return userStats, err
	}
	return userStats, nil
}

func (repo *UserStatsRepo) UpdatePlayerStats() ([]models.UserStats, error) {
	return nil, nil
}

func NewUserStatsRepo(db *gorm.DB) *UserStatsRepo {
	return &UserStatsRepo{db}
}
