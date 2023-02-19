package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type UserStatsController struct {
	repo *repos.UserStatsRepo
}

func (controller *UserStatsController) GetAllUserStats(c *fiber.Ctx) error {
	return nil
}

func (controller *UserStatsController) GetUserStats(c *fiber.Ctx) error {
	return nil
}

func NewUserStatsController(repo *repos.UserStatsRepo) *UserStatsController {
	return &UserStatsController{repo}
}

func RegisterUserStatsController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserStatsRepo(db)
	controller := NewUserStatsController(repo)

	UserStatsRouter := router.Group("/stats")

	UserStatsRouter.Get("/All", controller.GetAllUserStats)
	UserStatsRouter.Get("/user", controller.GetUserStats)

}
