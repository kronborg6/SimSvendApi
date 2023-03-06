package repos

import (
	"errors"
	"fmt"
	"time"

	"github.com/kronborg6/SimSvendApi/api/middleware"
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) Login(data models.UserInfo) (*[]models.User, error) {
	var user []models.User

	err := repo.db.Joins("Userinfo").Preload("UserStats").Preload("FriendList").Preload("Role").Find(&user, "email = ?", data.Email)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")

	}

	if !middleware.CheckPasswordHash(data.Password, user[0].Userinfo.Password) {
		return nil, errors.New("password not matchs")
	}
	user[0].Userinfo.Password = ""
	return &user, nil
}

// this is gona be check token endpoint pt just id
func (repo *UserRepo) CheckToken(id int) (*[]models.User, error) {
	var user []models.User

	err := repo.db.Joins("Userinfo").Preload("UserStats").Preload("FriendList").Preload("Role").Find(&user, "Userinfo.id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")

	}
	user[0].Userinfo.Password = ""
	return &user, nil
}
func (repo *UserRepo) FindUser(data models.UserInfo) ([]models.User, error) {
	var user []models.User
	err := repo.db.Joins("Userinfo").Preload("UserStats").Find(&user, "email = ?", data.Email)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")
	}
	user[0].Userinfo.Password = ""
	user[0].FriendList = nil

	return user, nil
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
func (repo *UserRepo) NewUser2(user models.UserInfo) (models.UserInfo, error) {

	user.CreateAt = time.Now()
	var userStats models.UserStats
	if err := repo.db.Create(&userStats).Error; err != nil {
		return user, err
	}
	if err := repo.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
func (repo *UserRepo) NewUser(user models.UserInfo) (models.User, error) {
	var newUser models.User
	var stats models.UserStats

	user.Password = middleware.HashPassword(user.Password)

	stats.Elo = 100
	newUser.UserStats = stats
	newUser.Userinfo = user
	newUser.Userinfo.CreateAt = time.Now()
	newUser.RoleID = 1
	// newUser.FriendList = nil

	if err := repo.db.Debug().Create(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (repo *UserRepo) FriendList(id int) ([]models.UserInfo, error) {
	var allFriends []models.UserInfo
	var friendslist []models.Friends

	err := repo.db.Debug().Where("user_id = ?", id).Find(&friendslist)
	if err.Error != nil {
		return nil, errors.New("user don't have friends LOL")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("user don't have friends LOL")
	}

	for i := range friendslist {
		var friends []models.UserInfo
		if err := repo.db.Where("email = ?", friendslist[i].Email).Find(&friends).Error; err != nil {
			return allFriends, err
		}
		friends[0].Password = ""
		allFriends = append(allFriends, friends...)
	}
	fmt.Println(id)
	return allFriends, nil
}

func (repo *UserRepo) FindAllPlayerStats() ([]models.UserStats, error) {
	var userStats []models.UserStats

	if err := repo.db.Debug().Find(&userStats).Error; err != nil {
		return userStats, err
	}
	return userStats, nil
}

func (repo *UserRepo) FindPlayerStats(id int) ([]models.User, error) {
	var user []models.User

	err := repo.db.Debug().Joins("UserStats").Find(&user, "UserStats.id = ?", 1)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")

	}
	return user, nil
}

func (repo *UserRepo) TopPlayers() ([]models.User, error) {

	var user []models.User

	if err := repo.db.Debug().Preload(clause.Associations).Find(&user).Error; err != nil {
		return nil, err
	}
	for i := range user {
		user[i].FriendList = nil
		user[i].Role.ID = 0
		user[i].Role.Name = ""
		user[i].Userinfo.Password = ""
		user[i].RoleID = 0
	}
	return user, nil
}

func (repo *UserRepo) UpdatePlayerStats() ([]models.UserStats, error) {
	return nil, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
