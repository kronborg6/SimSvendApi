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

	// match := middleware.CheckPasswordHash("test", "test")
	// fmt.Println("Match: ", match)

	controllers.RegisterAuthController(db, api)

	controllers.RegisterUserController(db, api)
	controllers.RegisterLeaderboardController(db, api)
	controllers.RegisterTournamentController(db, api)
	controllers.RegisterMatchController(db, api)
	// controllers.
	// controllers.RegisterUserStatsController(db, api)

	/* 	app.Get("/gtpp", func(c *fiber.Ctx) error {
	   		claims := jwt.MapClaims{
	   			"name":  "John Doe",
	   			"admin": true,
	   			"exp":   time.Now().Add(time.Hour * 72).Unix(),
	   		}

	   		// Create token
	   		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	   		// Generate encoded token and send it as response.
	   		t, err := token.SignedString([]byte(os.Getenv("PUBLIC")))
	   		if err != nil {
	   			return c.SendStatus(fiber.StatusInternalServerError)
	   		}

	   		return c.JSON(fiber.Map{"token": t})
	   	})
	   	app.Get("/gtp", func(c *fiber.Ctx) error {
	   		claims := jwt.MapClaims{
	   			"name": "Mikkel Kronborg",
	   			"exp":  time.Now().Add(time.Hour * 72).Unix(),
	   		}

	   		// Create token
	   		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	   		// Generate encoded token and send it as response.
	   		t, err := token.SignedString([]byte(os.Getenv("PUBLIC")))
	   		if err != nil {
	   			return c.SendStatus(fiber.StatusInternalServerError)
	   		}

	   		return c.JSON(fiber.Map{"token": t})
	   	})
	   	app.Get("/test", func(c *fiber.Ctx) error {
	   		return c.JSON(fiber.Map{
	   			"message": "this is a endpoint test",
	   			"test":    "Test",
	   		})
	   	})
	   	app.Use(jwtware.New(jwtware.Config{
	   		SigningKey: []byte("Mikkel"),
	   	}))
	   	app.Get("/ttp", func(c *fiber.Ctx) error {
	   		return c.JSON(fiber.Map{
	   			"message": "this is a endpoint test",
	   			"test":    "Test",
	   		})
	   	})

	   	app.Use(jwtware.New(jwtware.Config{
	   		SigningKey: []byte("Kronborg"),
	   	}))

	   	app.Get("/ttpp", func(c *fiber.Ctx) error {
	   		return c.JSON(fiber.Map{
	   			"message": "this is a endpoint test",
	   			"test":    "Test",
	   		})
	   	}) */

	log.Fatal(app.Listen(getPort()))
}
