package repo

import (
	"github.com/kronborg6/SimSvendApi/api/middleware"
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) FindUser(data models.UserInfo) (*[]models.UserInfo, error) {
	var user []models.UserInfo

	if err := repo.db.Where("email = ?", data.Email).Find(&user).Error; err != nil {
		return &user, err
	}

	if !middleware.CheckPasswordHash(user[0].Password, data.Password) {
		return nil, nil
	}

	return &user, nil
}

func (repo *UserRepo) FindAllUser() ([]models.UserInfo, error) {
	return nil, nil
}
func (repo *UserRepo) NewUser(user models.UserInfo) (*[]models.UserInfo, error) {
	return nil, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
