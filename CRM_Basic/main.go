package main

import (
	"CRM_Basic/database"
	"CRM_Basic/lead"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Define the routes and their corresponding handlers for a web application,
// built with the fiber web framework.
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead:id", lead.DeleteLead)
}

// Establish connection to the database.
func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("Failed to connect to DB.")
	}

	fmt.Println("Connected to DB.")
	database.DBConn.AutoMigrate((&lead.Lead{})) // Update the database schema for the Lead model.
	fmt.Println("Database migrated.")
}

func main() {
	app := fiber.New()            // Create a new instance of the fiber web framework.
	initDatabase()                // Establish connection to DB.
	setupRoutes(app)              // Define routes and handles for the server.
	app.Listen(3000)              // Start the server at port 3000.
	defer database.DBConn.Close() // Close the DB connection.
}
