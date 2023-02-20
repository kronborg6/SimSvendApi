package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/models"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type UserController struct {
	repo *repos.UserRepo
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {
	var user models.UserInfo
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.FindUser(user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())

		// panic("This panic is caught by fiber")
		// return c.JSON(fiber.Map{
		// 	"message": "username or password do not match",
		// 	"error":   err,
		// })
	}
	// if test, err = controller.repo.FindUser(user); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	// fmt.Print(data)
	return c.JSON(data)
}
func (controller *UserController) GetAllUser(c *fiber.Ctx) error {
	return nil
}
func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.UserInfo
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())

	}

	data, err := controller.repo.NewUser(user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(data)
}

func NewUserController(repo *repos.UserRepo) *UserController {
	return &UserController{repo}
}

func RegisterUserController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewUserController(repo)

	UserRouter := router.Group("/user")

	UserRouter.Post("/login", controller.GetUser)
	UserRouter.Post("/register", controller.CreateUser)
}
