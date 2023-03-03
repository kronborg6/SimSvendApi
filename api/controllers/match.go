package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type MatchController struct {
	repo *repos.MatchRepo
}

func (controller *MatchController) FindAllUserMatchs(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.GameHistory(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func NewMatchController(repo *repos.MatchRepo) *MatchController {
	return &MatchController{repo}
}

func RegisterMatchController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewMatchRepo(db)
	controller := NewMatchController(repo)

	MatchController := router.Group("/match")

	MatchController.Get("/:id", controller.FindAllUserMatchs)
}
