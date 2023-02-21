package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type LeaderboardController struct {
	repo *repos.LeaderboardRepo
}

func (controller *LeaderboardController) GetAllLeaderboard(c *fiber.Ctx) error {
	// var err error

	leaderboard, err := controller.repo.FindAllLeaderboards()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(leaderboard)

}

func NewLeaderbaordController(repo *repos.LeaderboardRepo) *LeaderboardController {
	return &LeaderboardController{repo}
}

func RegisterLeaderboardController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewLeaderbaordRepo(db)
	controller := NewLeaderbaordController(repo)

	LeaderbaordController := router.Group("/leaderbaord")

	LeaderbaordController.Get("/all", controller.GetAllLeaderboard)
}
