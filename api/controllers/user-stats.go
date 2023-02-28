package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type UserStatsController struct {
	repo *repos.UserStatsRepo
}

func (controller *UserStatsController) GetAllUserStats(c *fiber.Ctx) error {

	userStats, err := controller.repo.FindAllPlayerStats()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(userStats)
}

func (controller *UserStatsController) GetUserStats(c *fiber.Ctx) error {
	var err error
	// var email string
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data, err := controller.repo.FindPlayerStats(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(data)
}

func NewUserStatsController(repo *repos.UserStatsRepo) *UserStatsController {
	return &UserStatsController{repo}
}

func RegisterUserStatsController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserStatsRepo(db)
	controller := NewUserStatsController(repo)

	UserStatsRouter := router.Group("/stats")

	UserStatsRouter.Get("/All", controller.GetAllUserStats)
	UserStatsRouter.Get("/:id", controller.GetUserStats)

}
