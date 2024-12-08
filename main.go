package main

import (
	"log"

	"surya-sanity/go-todo/database"
	"surya-sanity/go-todo/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer func() {
		postgresDB, _ := database.DB.DB()
		postgresDB.Close()
	}()

	api := app.Group("/api")
	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":8080"))
}
