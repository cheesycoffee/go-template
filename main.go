package main

import (
	"fmt"

	// auto load .env file
	_ "github.com/joho/godotenv/autoload"

	"github.com/cheesycoffee/go-template/config"
)

func main() {
	cfg := config.New()
	fmt.Println(cfg.Env().Server)
}
