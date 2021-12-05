package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})
	app.Static("/public", "./public")
	app.Get("/", mainPage)
	app.Listen(":3000")
}

func mainPage(c *fiber.Ctx) {
	c.Render("mainpage", nil)
}
