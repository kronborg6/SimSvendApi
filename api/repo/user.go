package repo

import "gorm.io/gorm"

type UserRepo struct {
	db *gorm.DB
}