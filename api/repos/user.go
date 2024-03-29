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

	err := repo.db.Joins("Userinfo").Preload("UserStats").Preload("FriendList", "is_friends = true").Preload("Role").Find(&user, "email = ?", data.Email)
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
func (repo *UserRepo) UpdateStats(data models.UserStats) (models.UserStats, error) {
	// var stats models.UserStats

	// if err := repo.db.Debug().Save(&data).Error; err != nil {
	// 	return stats, nil
	// }
	err := repo.db.Debug().Model(&data).Where("user_id = ?", data.UserID).Updates(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
func (repo *UserRepo) UpdateUserRole(user models.User) (models.User, error) {

	err := repo.db.Debug().Model(&user).Where("id = ?", user.ID).Updates(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
func (repo *UserRepo) NewUser(user models.UserInfo) (models.User, error) {
	var newUser models.User
	var checkEmail models.UserInfo
	var stats models.UserStats

	user.Password = middleware.HashPassword(user.Password)

	stats.Elo = 100
	newUser.UserStats = stats
	newUser.Userinfo = user
	newUser.Userinfo.CreateAt = time.Now()
	newUser.RoleID = 1

	if err := repo.db.Debug().Where("email = ?", user.Email).Find(&checkEmail).Error; err != nil {
		return newUser, err
	}
	if checkEmail.Email == user.Email {
		return newUser, errors.New("email all ready exits")
	}
	if err := repo.db.Debug().Create(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (repo *UserRepo) FriendList(id int) ([]models.Friends, error) {
	var friendslist []models.Friends

	err := repo.db.Debug().Where("user_id = ? AND is_friends = true", id).Preload("Friend." + clause.Associations).Find(&friendslist)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("user don't have friends LOL")
	}
	for i := range friendslist {
		friendslist[i].Friend.FriendList = nil
		friendslist[i].Friend.RoleID = 0
		friendslist[i].Friend.Role.ID = 0
		friendslist[i].Friend.Role.Name = ""
		friendslist[i].Friend.Userinfo.Password = ""
	}

	return friendslist, nil
}
func (repo *UserRepo) FriendRequests(id int) ([]models.Friends, error) {
	var friendslist []models.Friends

	err := repo.db.Debug().Where("user_id = ? AND is_friends = false", id).Preload("Friend").Find(&friendslist)
	if err.Error != nil {
		return nil, errors.New("failed to find freinds <3")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("user don't have friends LOL")
	}

	fmt.Println(id)
	return friendslist, nil
}
func (repo *UserRepo) AcceptFriendRequest(freind models.Friends) (models.Friends, error) {
	var freindother models.Friends
	if err := repo.db.Debug().Where("user_id = ? AND friend_id = ? AND is_friends = false", freind.UserID, freind.FriendID).Preload("Friend").Find(&freind).Error; err != nil {
		return freind, err
	}
	if freind.Sender {
		return freindother, errors.New("you can accept a friend request you send")
	}
	fmt.Println(freind)
	freind.IsFriends = true
	freindother.UserID = freind.FriendID
	freindother.FriendID = freind.UserID
	freindother.IsFriends = true
	if err := repo.db.Debug().Model(&freind).Where("user_id = ? AND friend_id = ? AND is_friends = false", freind.UserID, freind.FriendID).Updates(&freind).Error; err != nil {
		return freind, err
	}
	if err := repo.db.Debug().Model(&freindother).Where("user_id = ? AND friend_id = ? AND is_friends = false", freindother.UserID, freindother.FriendID).Updates(&freindother).Error; err != nil {
		return freind, err
	}
	return freind, nil
}
func (repo *UserRepo) SendFriendRequest(friend models.Friends) error {
	var checkfriend models.Friends
	var ortherfriend models.Friends
	if err := repo.db.Debug().Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", friend.UserID, friend.FriendID, friend.FriendID, friend.UserID).First(&checkfriend).Error; err == nil {
		return fmt.Errorf("users %d and %d are already friends", friend.UserID, friend.FriendID)
	}

	ortherfriend.FriendID = friend.UserID
	ortherfriend.UserID = friend.FriendID
	friend.Sender = true
	if err := repo.db.Debug().Create(&friend).Error; err != nil {
		return err
	}
	if err := repo.db.Debug().Create(&ortherfriend).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) RemoveFriendAndRequest(freind models.Friends) error {

	err := repo.db.Debug().Where("user_id = ? AND friend_id = ?", freind.UserID, freind.FriendID).Delete(&freind)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return errors.New("user don't have friends LOL")
	}
	er := repo.db.Debug().Where("user_id = ? AND friend_id = ?", freind.FriendID, freind.UserID).Delete(&freind)
	if er.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return errors.New("user don't have friends LOL")
	}
	return nil
}

func (repo *UserRepo) DenieFriendRequest(friend models.Friends) error {
	return nil
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
