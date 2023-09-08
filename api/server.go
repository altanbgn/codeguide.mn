package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"

  "github.com/gofiber/fiber/v2"

  "github.com/altanbgn/arctis/api/config"
  "github.com/altanbgn/arctis/api/middleware"
  "github.com/altanbgn/arctis/api/modules"
)

func Run() {
  appConfig := config.AppConfig()
  _ = config.DBConfig()

  // TODO Connect to database
  fiberConfig := config.FiberConfig()
  app := fiber.New(fiberConfig)

  // Middlewares
  middleware.FiberMiddleware(app)

  // Routes


  // Graceful shutdown
  sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

  go func () {
    <- sigChan
    fmt.Println("Gracefully putting server to sleep...")
    _ = app.Shutdown()
  }()

  serverAddr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)

  if error := app.Listen(serverAddr); error != nil {
    fmt.Errorf("Server not runnning! error: %v", error)
  }
}
