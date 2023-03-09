package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/kronborg6/SimSvendApi/api/controllers"
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
	// app.Use(cache.New())
	app.Use(cache.New(cache.Config{

		Expiration:   5 * time.Second,
		CacheControl: true,
	}))
	// middleware.EncryptoKey()
	// fmt.Println(getPort())

	api := app.Group("/")

	controllers.RegisterAuthController(db, api)
	controllers.RegisterUserController(db, api)
	controllers.RegisterLeaderboardController(db, api)
	controllers.RegisterTournamentController(db, api)
	controllers.RegisterMatchController(db, api)

	log.Fatal(app.Listen(getPort()))
}
