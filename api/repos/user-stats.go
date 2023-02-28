package repos

import (
	"errors"

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

func (repo *UserStatsRepo) FindPlayerStats(id int) ([]models.User, error) {
	var user []models.User
	// var userinfo []models.UserInfo

	err := repo.db.Debug().Joins("UserStats").Find(&user, "UserStats.id = ?", 1)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")

	}
	// user[0].Userinfo =

	// if err := repo.db.Debug().Where("id = ?", userinfo[0].UserStatsID).Find(userStats).Error; err != nil {
	// 	return userStats, err
	// }
	return user, nil
}

func (repo *UserStatsRepo) UpdatePlayerStats() ([]models.UserStats, error) {
	return nil, nil
}

func NewUserStatsRepo(db *gorm.DB) *UserStatsRepo {
	return &UserStatsRepo{db}
}
