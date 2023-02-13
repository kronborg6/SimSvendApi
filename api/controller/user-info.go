package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/SimSvendApi/api/repo"
)

type UserController struct {
	repo *repo.UserRepo
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {
	return nil
}
func (controller *UserController) GetAllUser(c *fiber.Ctx) error {
	return nil
}
func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	return nil
}
