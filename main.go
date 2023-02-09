package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kronborg6/SimSvendApi/api/db"
	"github.com/kronborg6/SimSvendApi/api/models"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port

}

func main() {
	db := db.Init()
	app := fiber.New()
	models.Setup(db)

	app.Use(logger.New())
	// fmt.Println(getPort())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this is a endpoint test",
			"test":    "Test",
		})
	})

	log.Fatal(app.Listen(getPort()))
}
