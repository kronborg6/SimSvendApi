package repos

import (
	"errors"
	"fmt"
	"time"

	"github.com/kronborg6/SimSvendApi/api/middleware"
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) FindUser(data models.UserInfo) (*[]models.User, error) {
	var user []models.User
	fmt.Println("lol")

	// if err := repo.db.Find(&user).Error; err != nil {
	// 	fmt.Println("Hej med dig")
	// 	return nil, err
	// }
	err := repo.db.Where("email = ?", data.Email).Joins("UserInfo").Preload("UserInfo").Preload("UserStats").Find(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")

	}
	fmt.Println(user[0])

	// fmt.Println(user[0].Password)
	if !middleware.CheckPasswordHash(data.Password, user[0].Userinfo.Password) {
		return nil, errors.New("password not matchs")
	}
	user[0].Userinfo.Password = ""
	return &user, nil
}

func (repo *UserRepo) FindAllUser() ([]models.User, error) {
	var user []models.User

	err := repo.db.Preload("FriendList").Where("id = 3").Find(&user)

	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")
	}
	user[0].Userinfo.Password = ""
	return user, nil
}
func (repo *UserRepo) NewUser(user models.UserInfo) (models.UserInfo, error) {

	user.CreateAt = time.Now()
	var userStats models.UserStats
	if err := repo.db.Create(&userStats).Error; err != nil {
		return user, err
	}
	// user.UserStatsID = userStats.Id
	if err := repo.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepo) FriendList() ([]models.UserInfo, error) {
	return nil, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
