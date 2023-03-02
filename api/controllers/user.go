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
func (controller *UserController) CheckToken(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.CheckToken(id)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
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
func (controller *UserController) CreateNewUser(c *fiber.Ctx) error {
	var model models.UserInfo
	var err error
	if err = c.BodyParser(&model); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.NewUser(model)
	if err != nil {
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

	data, err := controller.repo.NewUser2(user)

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

func (controller *UserController) GetAllUserStats(c *fiber.Ctx) error {

	userStats, err := controller.repo.FindAllPlayerStats()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(userStats)
}

func (controller *UserController) GetUserStats(c *fiber.Ctx) error {
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

func (controller *UserController) GetTopPlayers(c *fiber.Ctx) error {
	data, err := controller.repo.TopPlayers()

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

	UserRouter.Post("/login", controller.Login)
	UserRouter.Post("/register", controller.CreateUser)
	UserRouter.Get("/test", controller.GetAllUser)
	UserRouter.Post("/find", controller.GetUser)

	UserRouter.Post("/", controller.CreateNewUser)

	UserRouter.Get("/token/:id", controller.CheckToken)

	UserRouter.Get("/friends/:userID", controller.TestFriends)

	UserRouter.Get("/leaderboard", controller.GetTopPlayers)

	UserStatsRouter := router.Group("/stats")

	UserStatsRouter.Get("/All", controller.GetAllUserStats)
	UserStatsRouter.Get("/:id", controller.GetUserStats)
}
