package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/models"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type CasualController struct {
	repo *repos.CasualRepo
}

func (controller *CasualController) CreateCasualGame(c *fiber.Ctx) error {
	var match models.Match
	var err error
	if err = c.BodyParser(&match); err != nil {
		// return c.JSON(match)
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}
	match.PlayTime = time.Now()
	data, err := controller.repo.NewCasualGame(match)

	if err != nil {
		// return c.JSON(match)

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(data)
}

func (controller *CasualController) PutMatchResult(c *fiber.Ctx) error {
	var result models.Results
	var err error

	if err = c.BodyParser(&result); err != nil {
		return c.JSON(result)
	}
	data, err := controller.repo.AddResualtCasualGame(result)

	if err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}
	return c.JSON(data)
}

func NewCasualController(repo *repos.CasualRepo) *CasualController {
	return &CasualController{repo}
}

func RegisterCasualController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewCasualRepo(db)
	controller := NewCasualController(repo)

	CasualController := router.Group("/casual")

	CasualController.Post("/", controller.CreateCasualGame)
	CasualController.Post("/result", controller.PutMatchResult)
}
