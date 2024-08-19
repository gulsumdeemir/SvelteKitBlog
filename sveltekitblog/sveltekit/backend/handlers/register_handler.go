package handlers

import (
	"hello-world/backend/database"
	"hello-world/backend/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz veri"})
    }

    if result := database.DB.Create(&user); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı ekleme hatası"})
    }
    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Kullanıcı başarıyla oluşturuldu"})
}
