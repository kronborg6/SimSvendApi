package repos

import (
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type UserStatsRepo struct {
	repo *gorm.DB
}

func (repo *UserStatsRepo) FindAllPlayerStats() (*[]models.UserStats, error) {
	return nil, nil
}

func (repo *UserStatsRepo) FindPlayerStats(email string) ([]models.UserStats, error) {
	return nil, nil
}

func (repo *UserStatsRepo) UpdatePlayerStats() ([]models.UserStats, error) {
	return nil, nil
}
