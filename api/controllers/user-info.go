package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/models"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type UserController struct {
	repo *repos.UserRepo
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var user models.UserInfo
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.Login(user)

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
func (controller *UserController) GetUser(c *fiber.Ctx) error {
	var user models.UserInfo
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println(user)
	data, err := controller.repo.FindUser(user)
	if err != nil {
		fmt.Println("hej")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}
func (controller *UserController) GetAllUser(c *fiber.Ctx) error {
	// var user models.User
	// var err error

	data, err := controller.repo.FindAllUser()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())

		// panic("This panic is caught by fiber")
		// return c.JSON(fiber.Map{
		// 	"message": "username or password do not match",
		// 	"error":   err,
		// })
	}

	return c.JSON(data)
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

func (controller *UserController) TestFriends(c *fiber.Ctx) error {
	var err error

	id, err := strconv.Atoi(c.Params("userID"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data, err := controller.repo.FriendList(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// if data == nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())

	// }

	return c.JSON(data)
}

func NewUserController(repo *repos.UserRepo) *UserController {
	return &UserController{repo}
}

func RegisterUserController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewUserController(repo)

	UserRouter := router.Group("/user")

	UserRouter.Post("/login", controller.Login)
	UserRouter.Post("/register", controller.CreateUser)
	UserRouter.Get("/test", controller.GetAllUser)
	UserRouter.Post("/find", controller.GetUser)

	UserRouter.Get("/friends/:userID", controller.TestFriends)
}
