package handlers

import (
	"errors"
	"hello-world/backend/database"
	"hello-world/backend/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)


func LoginHandler(c *fiber.Ctx) error {

	if database.DB == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Veritabanı bağlantısı sağlanamadı"})
	}

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz giriş bilgileri"})
	}

	var user models.User
	if err := database.DB.Where("eMail = ?", body.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz e-posta veya şifre"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Bir hata oluştu"})
	}

	if user.UserPassword != body.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz e-posta veya şifre"})
	}

	token, err := GenerateToken(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token oluşturulamadı"})
	}

	return c.JSON(fiber.Map{"token": token, "userName": user.UserName, "email": user.EMail, "userID": user.UserID})
}



func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "", errors.New("JWT_SECRET_KEY çevresel değişkeni ayarlanmamış")
	}
	return token.SignedString([]byte(secret))
}
