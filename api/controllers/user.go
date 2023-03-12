package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(data)
}
func (controller *UserController) CheckToken(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.CheckToken(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(data)
}
func (controller *UserController) GetFreindList(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.FriendList(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(data)
}
func (controller *UserController) GetAllUser(c *fiber.Ctx) error {
	// var user models.User
	// var err error

	data, err := controller.repo.FindAllUser()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())

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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(data)
}

func (controller *UserController) PutUserStats(c *fiber.Ctx) error {
	var stats models.UserStats
	var err error
	if err = c.BodyParser(&stats); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.UpdateStats(stats)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(data)
}
func (controller *UserController) PutUserRole(c *fiber.Ctx) error {
	var user models.User
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.UpdateUserRole(user)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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

func (controller *UserController) PostAcceptFriend(c *fiber.Ctx) error {
	var freind models.Friends
	var err error

	if err = c.BodyParser(&freind); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("controller")
	fmt.Println(freind)
	data, err := controller.repo.AcceptFriendRequest(freind)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(data)
}
func (controller *UserController) PostRemoveFriend(c *fiber.Ctx) error {
	var freind models.Friends
	var err error

	if err = c.BodyParser(&freind); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("controller")
	fmt.Println(freind)
	// _, err := controller.repo.RemoveFriendAndRequest(freind)
	err = controller.repo.RemoveFriendAndRequest(freind)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"delete": "Fuck that friend he is gone",
	})
}
func (controller *UserController) PostFriendRequest(c *fiber.Ctx) error {
	var freind models.Friends
	var err error

	if err = c.BodyParser(&freind); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	err = controller.repo.SendFriendRequest(freind)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"request": "friend request is send",
	})
}
func NewUserController(repo *repos.UserRepo) *UserController {
	return &UserController{repo}
}

func RegisterUserController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewUserController(repo)

	FreindsRouter := router.Group("/freinds")

	FreindsRouter.Get("/:id", controller.GetFreindList)
	FreindsRouter.Post("/accept", controller.PostAcceptFriend)
	FreindsRouter.Post("/new", controller.PostFriendRequest)
	FreindsRouter.Post("/remove", controller.PostRemoveFriend)
	FreindsRouter.Get("/:userID", controller.TestFriends)

	UserRouter := router.Group("/user")

	UserRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PUBLIC")),
	}))

	// UserRouter.Post("/login", controller.Login)
	// UserRouter.Post("/register", controller.CreateUser)
	UserRouter.Get("/test", controller.GetAllUser)
	UserRouter.Post("/find", controller.GetUser)

	// UserRouter.Post("/", controller.CreateNewUser)

	UserRouter.Get("/token/:id", controller.CheckToken)

	UserRouter.Get("/leaderboard", controller.GetTopPlayers)

	AdminRouter := router.Group("/admin")

	//admin side endpoint
	AdminRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PRIVATE")),
	}))
	AdminRouter.Post("/stats", controller.PutUserStats)
	AdminRouter.Post("/role", controller.PutUserRole)
	AdminRouter.Get("/all", controller.GetTopPlayers)

	UserStatsRouter := router.Group("/stats")
	UserStatsRouter.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("PUBLIC")),
	}))

	UserStatsRouter.Get("/All", controller.GetAllUserStats)
	UserStatsRouter.Get("/:id", controller.GetUserStats)
}
