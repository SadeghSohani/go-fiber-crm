package models

import (
	"github.com/SadeghSohani/go-fiber-crm/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.Db
	var leads []Lead
	db.Find(&leads)
	err := c.JSON(leads)
	if err != nil {
		return err
	}
	return nil
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Db
	var lead Lead
	db.Find(&lead, id)
	err := c.JSON(lead)
	if err != nil {
		return err
	}
	return nil
}

func NewLead(c *fiber.Ctx) error {
	db := database.Db
	var lead Lead
	if err := c.BodyParser(&lead); err != nil {
		err := c.Status(503).JSON(err)
		if err != nil {
			return err
		}
		return err
	}
	db.Create(&lead)
	//c.JSON(lead)
	err := c.JSON(lead)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Db
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		err := c.Status(500).SendString("No lead with ID")
		if err != nil {
			return err
		}
	}
	db.Delete(&lead)
	err := c.Status(200).SendString("Lead deleted successfully.")
	if err != nil {
		return err
	}
	return nil
}
