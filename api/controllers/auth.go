package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kronborg6/SimSvendApi/api/models"
	"github.com/kronborg6/SimSvendApi/api/repos"
	"gorm.io/gorm"
)

type AuthController struct {
	repo *repos.UserRepo
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	var user models.UserInfo
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.Login(user)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	claims := jwt.MapClaims{
		"user": (*data)[0].Userinfo,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("PUBLIC")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "user": data})
}
func (controller *AuthController) AdminLogin(c *fiber.Ctx) error {
	var user models.UserInfo
	var err error

	if err = c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.Login(user)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	claims := jwt.MapClaims{
		"user": (*data)[0].Userinfo,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("PRIVATE")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "user": data})
}
func (controller *AuthController) CheckToken(c *fiber.Ctx) error {
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
func (controller *AuthController) CreateNewUser(c *fiber.Ctx) error {
	var model models.UserInfo
	var err error
	if err = c.BodyParser(&model); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	data, err := controller.repo.NewUser(model)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	claims := jwt.MapClaims{
		"user": data.Userinfo,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("PUBLIC")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "user": data})
}
func NewAuthController(repo *repos.UserRepo) *AuthController {
	return &AuthController{repo}
}

func RegisterAuthController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewAuthController(repo)

	AuthRouter := router.Group("/auth")

	AuthRouter.Post("/login", controller.Login)
	AuthRouter.Post("/adminlogin", controller.AdminLogin)
	AuthRouter.Post("/register", controller.CreateNewUser)
	AuthRouter.Get("/:id", controller.CheckToken)
}
