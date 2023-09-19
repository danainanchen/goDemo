package main

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Fname string
	Lname string
}

func getPerson(ctx *fiber.Ctx) error {
	person_1 := Person{
		Fname: "Danainan",
		Lname: "Chamnanpaison",
	}
	return ctx.Status(fiber.StatusOK).JSON(person_1)
}

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})
	app.Get("/person", getPerson)
	app.Listen(":80")
}
