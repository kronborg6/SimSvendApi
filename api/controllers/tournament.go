package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/models"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type TournamentController struct {
	repo *repos.TournamentRepo
}

func (controller *TournamentController) GetAllTour(c *fiber.Ctx) error {

	tour, err := controller.repo.FindAllTour()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(tour)
}

func (controller *TournamentController) GetTour(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	tour, err := controller.repo.FindTour(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(tour)
}
func (controller *TournamentController) CreateNewTour(c *fiber.Ctx) error {
	var tour models.Tournament
	var err error

	if err = c.BodyParser(&tour); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.NewTour(tour)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func NewTournamentController(repo *repos.TournamentRepo) *TournamentController {
	return &TournamentController{repo}
}

func RegisterTournamentController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewTournamentRepo(db)
	controller := NewTournamentController(repo)

	TourController := router.Group("/tour")

	TourController.Get("/", controller.GetAllTour)
	TourController.Get("/:id", controller.GetTour)
	TourController.Post("/", controller.CreateNewTour)

}
