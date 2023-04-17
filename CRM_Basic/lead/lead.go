package lead

import (
	"CRM_Basic/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

// Get the leads info from DB.
func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

// Get the lead info from DB.
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead

	db.Find(&lead, id)
	c.JSON(lead)
}

// Add a new lead in the DB.
func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)

	// Check that the info sent is correct.
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	c.JSON(lead)
}

// Delete a lead from DB.
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(500).Send("No lead found with this ID: %s.", id)
		return
	}

	db.Delete(&lead)
	c.Send("Lead deleted from DB.")
}
