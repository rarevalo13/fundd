package fiberserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var app = fiber.New()

func Server() {
	app.Static("/", "./frontend/build")

	app.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./frontend/build/index.html")
	})

	SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

func SetupRoutes(app *fiber.App) {
	app.Post("api/v1/fundd/token", GetToken)
	app.Post("api/v1/fundd/accesstoken", GetAccessToken)
	app.Post("api/v1/fundd/recent", GetRecentTxns)
}
