package middleware

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			fmt.Println("Hata: Yetkilendirme başlığı eksik")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkilendirme gerekiyor"})
		}

		// Token'ı işleme
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

		claims, err := validateJWT(token)
		if err != nil {
			fmt.Println("Hata: JWT doğrulama hatası:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz veya süresi dolmuş token"})
		}

		
		userID, ok := claims["userID"].(float64)
		if !ok {
			fmt.Println("Hata: JWT claims hatası: userID bulunamadı")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Geçersiz token claims"})
		}

		fmt.Println("Başarılı JWT doğrulaması, userID:", userID)
		c.Locals("userID", int(userID))

		return c.Next()
	}
}

func validateJWT(token string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET_KEY") // Gizli anahtarı çevresel değişkenden al
	if secret == "" {
		fmt.Println("Hata: JWT_SECRET_KEY çevresel değişkeni ayarlanmamış")
		return nil, errors.New("gizli anahtar mevcut değil")
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("beklenmeyen imza yöntemi")
		}
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("JWT parse hatası:", err)
		return nil, err
	}

	if !parsedToken.Valid {
		fmt.Println("Geçersiz token")
		return nil, errors.New("geçersiz token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Claims hatası")
		return nil, errors.New("geçersiz token claims")
	}

	return claims, nil
}
