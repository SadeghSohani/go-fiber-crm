package main

import (
	"fmt"
	"github.com/SadeghSohani/go-fiber-crm/database"
	"github.com/SadeghSohani/go-fiber-crm/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", models.GetLeads)
	app.Get("/api/v1/lead/:id", models.GetLead)
	app.Post("/api/v1/lead", models.NewLead)
	app.Delete("/api/v1/lead/:id", models.DeleteLead)
}

func initDatabase() {
	var err error
	database.Db, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database.")
	}
	fmt.Println("Connection opened to database.")
	database.Db.AutoMigrate(&models.Lead{})
	fmt.Println("database migrated.")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	if err := app.Listen("localhost:3000"); err != nil {
		return
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(database.Db)
}
