package controllers

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"

	jwtware "github.com/gofiber/jwt/v3"
)

type LeaderboardController struct {
	repo *repos.LeaderboardRepo
}

func (controller *LeaderboardController) GetAllLeaderboard(c *fiber.Ctx) error {
	var err error
	// if err = c.BodyParser(&mounthNumber); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	leaderboard, err := controller.repo.FindAllLeaderboards(int(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(leaderboard)

}

func NewLeaderbaordController(repo *repos.LeaderboardRepo) *LeaderboardController {
	return &LeaderboardController{repo}
}

func RegisterLeaderboardController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewLeaderbaordRepo(db)
	controller := NewLeaderbaordController(repo)

	LeaderbaordController := router.Group("/leaderboard")
	LeaderbaordController.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PUBLIC")),
	}))
	LeaderbaordController.Get("/:id", controller.GetAllLeaderboard)
}
