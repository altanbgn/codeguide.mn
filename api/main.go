package main

import (
	"github.com/altanbgn/arctis/api/config"
)

func main() {
  config.LoadEnvFile(".env")
  config.LoadDBConfig()
  config.LoadApp()

  Run()
}
