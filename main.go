package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Person struct {
	Fname string
	Lname string
}

var person Person

func getPerson(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(person)
}

func postPerson(ctx *fiber.Ctx) error {
	body := new(Person)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
		return err
	}

	person = Person{
		Fname: body.Fname,
		Lname: body.Lname,
	}
	return ctx.Status(fiber.StatusOK).JSON(person)
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello World")
	})

	app.Use(logger.New())
	app.Use(requestid.New())

	personApp := app.Group("/person")
	personApp.Get("", getPerson)
	personApp.Post("", postPerson)
	personApp.Get("/:fname/:lname", func(ctx *fiber.Ctx) error {
		firstName := ctx.Params("fname")
		lastName := ctx.Params("lname")
		res := "firstname: " + firstName + "\nlastname: " + lastName
		return ctx.SendString(res)
	})

	app.Listen(":8080")
}
