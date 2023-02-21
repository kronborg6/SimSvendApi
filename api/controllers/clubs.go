package controllers

import (
	"github.com/kronborg6/SimSvendApi/api/repos"
)

type ClubsController struct {
	repo *repos.ClubsRepo
}

// func (controller *ClubsController) GetAllClubs(c fiber.Ctx) error
// func (controller *ClubsController) GetClubs(c fiber.Ctx) error
// func (controller *ClubsController) PostClubs(c fiber.Ctx) error
// func (controller *ClubsController) PutClubs(c fiber.Ctx) error
// func (controller *ClubsController) DeleteClubs(c fiber.Ctx) error

func NewClubsRepo(repo *repos.ClubsRepo) *ClubsController {
	return &ClubsController{repo}
}
