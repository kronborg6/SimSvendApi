package controllers

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/kronborg6/SimSvendApi/api/models"
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

func (controller *MatchController) FindGame(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data, err := controller.repo.FindGame(id)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func (controller *MatchController) PutMatchResult(c *fiber.Ctx) error {
	var result models.Results
	var err error

	if err = c.BodyParser(&result); err != nil {
		return c.JSON(result)
	}
	data, err := controller.repo.SetGameScore(result)

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
	MatchController.Post("/setresult", controller.PutMatchResult)

	MatchController.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PUBLIC")),
	}))

	MatchController.Get("/:id", controller.FindAllUserMatchs)
	MatchController.Get("/game/:id", controller.FindGame)
}
