package handlers

import (
	"hello-world/backend/database"
	"hello-world/backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategory(c *fiber.Ctx) error {
    var categories []models.Category
    result := database.DB.Find(&categories)
    if result.Error != nil {
        return c.Status(500).SendString("Kategoriler alınamadı: " + result.Error.Error())
    }
    return c.JSON(categories)
}




func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	result := database.DB.Create(&category)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.Status(201).JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category


	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kategori bulunamadı")
	}


	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&category)

	return c.Status(200).JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category

	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kategori Bulunamadı")
	}
	database.DB.Delete(&category)

	return c.Status(200).SendString("Kategori Silindi")
}