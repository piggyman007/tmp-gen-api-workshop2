package main

import (
	"log"
	"os"
	"workshop2/model"
	"workshop2/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	if _, err := os.Stat("users.db"); os.IsNotExist(err) {
		os.Create("users.db")
	}
	// Connect to SQLite
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.User{}, &model.Transfer{})

	app := fiber.New()
	app.Use(logger.New())
	route.RegisterRoutes(app, db)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
