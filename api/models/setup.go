package models

import "gorm.io/gorm"

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Clubs{},
	)
}
