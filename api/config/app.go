package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Host        string
	Port        int
	Debug       bool
	ReadTimeout time.Duration

	JWTSecretKey    string
	JWTSecretExpire int
}

var app = &App{}

func AppConfig() *App {
  return app
}

func LoadApp() {
  app.Host = os.Getenv("APP_HOST")
  app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
  app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
  timeout, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
  app.ReadTimeout = time.Duration(timeout) * time.Second

  app.JWTSecretKey = os.Getenv("APP_SECRET_KEY")
  app.JWTSecretExpire, _ = strconv.Atoi(os.Getenv("JWT_SECRET_EXPIRE"))
}
