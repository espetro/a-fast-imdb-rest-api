package main

import (
	"datapi/config"
	"datapi/database"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/film", database.GetFilm)
	app.Get("/film/random", database.GetRandomFilm)
	app.Get("/player", database.GetPlayer)
	app.Get("/player/random", database.GetRandomPlayer)
}

func main() {
	defer database.Close()

	app := fiber.New()
	setupRoutes(app)
	app.Listen(config.APIPort)
}