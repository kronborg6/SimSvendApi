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

type TournamentController struct {
	repo *repos.TournamentRepo
}

func (controller *TournamentController) GetAllTour(c *fiber.Ctx) error {

	tour, err := controller.repo.FindAllTour()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(data)
}

func (controller *TournamentController) PostJoinTour(c *fiber.Ctx) error {
	var user models.JoinTourModel
	var err error
	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.JoinTour(user)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(data)
}

func NewTournamentController(repo *repos.TournamentRepo) *TournamentController {
	return &TournamentController{repo}
}

func (controller *TournamentController) PutTour(c *fiber.Ctx) error {
	var tour models.Tournament
	var err error
	if err = c.BodyParser(&tour); err != nil {
		return c.JSON(tour)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err = controller.repo.UpdateTour(tour); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(fiber.Map{
		"tournament": "updated",
	})
}

func (controller *TournamentController) PutTourInfo(c *fiber.Ctx) error {
	var info models.TournamentInfo
	var err error
	if err = c.BodyParser(&info); err != nil {
		return c.JSON(info)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err = controller.repo.UpdateTourInfo(info); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(fiber.Map{
		"tournament_info": "updated",
	})
}

func RegisterTournamentController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewTournamentRepo(db)
	controller := NewTournamentController(repo)

	TourController := router.Group("/tour")

	TourController.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PUBLIC")),
	}))

	TourController.Get("/", controller.GetAllTour)

	TourController.Get("/:id", controller.GetTour)
	TourController.Post("/join", controller.PostJoinTour)

	AdminTourController := router.Group("/admin/tour")

	AdminTourController.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PRIVATE")),
	}))
	AdminTourController.Get("/all", controller.GetAllTour)
	AdminTourController.Post("/", controller.CreateNewTour)
	AdminTourController.Post("/update", controller.PutTour)
	AdminTourController.Post("/:id/info", controller.PutTourInfo)

}
