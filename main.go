package main

import (
	"log"
	"nso-server/internal/app"
  _ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := app.Bootstrap(); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
