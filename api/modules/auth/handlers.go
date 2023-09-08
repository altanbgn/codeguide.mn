package auth

import "github.com/gofiber/fiber/v2"

func HandleLogin(ctx *fiber.Ctx) error {
  return ctx.JSON(fiber.Map {
    "hello": "login",
  })
}

func HandleRegister(ctx *fiber.Ctx) error {
  return ctx.JSON(fiber.Map{
    "hello": "register",
  })
}
