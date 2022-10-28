package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zazhedho/telkom-test/task6-api/src/config"
)

func main() {
	if err := config.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
