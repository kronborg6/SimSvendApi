package repos

import (
	"fmt"

	"github.com/kronborg6/SimSvendApi/api/middleware"
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) FindUser(data models.UserInfo) (*[]models.UserInfo, error) {
	var user []models.UserInfo
	fmt.Println("lol")

	if err := repo.db.Where("email = ?", data.Email).Find(&user).Error; err != nil {
		return &user, err
	}
	// fmt.Println(user[0].Password)
	if !middleware.CheckPasswordHash(data.Password, user[0].Password) {
		// if !middleware.CheckPasswordHash(user[0].Password, data.Password) {
		fmt.Println("password no match")
		return nil, nil
	}
	// fmt.Println(user)
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
