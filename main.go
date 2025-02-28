package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
)

func main() {
	fmt.Println("Hello, World! from file_linki")

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	app := fiber.New(fiber.Config{
		Prefork:      false,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 180 * time.Second,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())

	app.Get("/metrics", monitor.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"messege": "hello, world!",
		})
	})

	app.Listen(":8080")
}
