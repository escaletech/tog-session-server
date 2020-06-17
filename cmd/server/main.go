package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"

	"github.com/escaletech/tog-session-server/internal/config"
	"github.com/escaletech/tog-session-server/internal/handler"
)

func main() {
	logger := log.New(os.Stderr, "tog: ", log.LstdFlags|log.Lshortfile)
	godotenv.Load()

	conf, err := config.FromEnv()
	if err != nil {
		logger.Fatal(err)
	}

	var app *fiber.App
	retryLimit := time.Now().Add(5 * time.Minute)
	for {
		app = fiber.New(&fiber.Settings{
			DisableStartupMessage: true,
		})
		app.Use(middleware.Logger())
		err = handler.Register(app, conf)
		if err == nil || time.Now().After(retryLimit) {
			break
		}

		logger.Println(err)
		<-time.After(2 * time.Second)
	}

	if err != nil {
		logger.Fatal("retry limit exceeded, exiting")
	}

	addr := 3000
	log.New(os.Stdout, "tog: ", log.LstdFlags|log.Lshortfile).Println("server starting on", addr)
	if err := app.Listen(addr); err != nil {
		logger.Fatal(err)
	}
}
