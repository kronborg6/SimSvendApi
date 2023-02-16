package repos

import (
	"errors"
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

	// if err := repo.db.Find(&user).Error; err != nil {
	// 	fmt.Println("Hej med dig")
	// 	return nil, err
	// }
	err := repo.db.Where("email = ?", data.Email).Find(&user)
	fmt.Println("dsfgdfgdgfdgfdfg dfg dfg dfg df")
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find user")

	}
	fmt.Println(user[0].Password)

	// fmt.Println(user[0].Password)
	if !middleware.CheckPasswordHash(data.Password, user[0].Password) {
		fmt.Println("password no match")
		return nil, nil
	}
	fmt.Println("dsfgdfgdgfdgfdfg dfg dfg dfg df")

	return &user, nil
}

func (repo *UserRepo) FindAllUser() ([]models.UserInfo, error) {
	return nil, nil
}
func (repo *UserRepo) NewUser(user models.UserInfo) (models.UserInfo, error) {

	if err := repo.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
