package repos

import (
	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type ClubsRepo struct {
	db *gorm.DB
}

func (repo *ClubsRepo) FindAllClubs() ([]models.Clubs, error) { return nil, nil }

// func (repo *ClubsRepo) FindClub(id int) (models.Clubs, error)
// func (repo *ClubsRepo) NewClub(club models.Clubs) ([]models.Clubs, error)
// func (repo *ClubsRepo) EditClub(club models.Clubs) ([]models.Clubs, error)
// func (repo *ClubsRepo) DeleteClub(id int) ([]models.Clubs, error)

// func (repo *ClubsRepo) NewCourt(club models.ClubCourts) ([]models.ClubCourts, error)
// func (repo *ClubsRepo) EditCourt(club models.ClubCourts) ([]models.ClubCourts, error)
// func (repo *ClubsRepo) DeleteCourt(club models.ClubCourts) ([]models.ClubCourts, error)

func NewClubsRepo(db *gorm.DB) *ClubsRepo {
	return &ClubsRepo{db}
}