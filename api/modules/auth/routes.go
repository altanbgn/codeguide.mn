package auth

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
  route := app.Group("/v1/auth")

  route.Post("/login", HandleLogin)
  route.Post("/register", HandleRegister)
}
