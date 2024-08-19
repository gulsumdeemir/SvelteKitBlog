package handlers

import (
	"hello-world/backend/database"
	"hello-world/backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetSave(c *fiber.Ctx) error {
	var save []models.Save
	result := database.DB.Find(&save)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(save)
}


func CreateSave(c *fiber.Ctx) error {
	save := new(models.Save)
	if err := c.BodyParser(save); err != nil {
		log.Println("Body parser error:", err)
		return c.Status(400).SendString(err.Error())
	}

	if save.UserID == 0 { 
		return c.Status(400).SendString("Geçersiz kullanıcı ID'si")
	}

	var user models.User
	if err := database.DB.First(&user, save.UserID).Error; err != nil {
		return c.Status(400).SendString("Kullanıcı bulunamadı")
	}

	var post models.Post
	if err := database.DB.First(&post, save.PostID).Error; err != nil {
		return c.Status(400).SendString("Gönderi bulunamadı")
	}

	save.PostID = post.PostID;
	save.CategoryID = *post.CategoryID;
	

	result := database.DB.Create(&save)
	if result.Error != nil {
		log.Println("Database error:", result.Error)
		return c.Status(500).SendString(result.Error.Error())
	}

	return c.Status(201).JSON(save)
}

  
  
func UpdateSave(c *fiber.Ctx) error {
	id := c.Params("id")
	var save models.Save
	result := database.DB.First(&save, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kaydedilme bulunamadı")
	}
	if err := c.BodyParser(&save); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&save)
	return c.Status(200).JSON(save)
}


func DeleteSave(c *fiber.Ctx) error {
	id := c.Params("id")
	var save models.Save
	result := database.DB.First(&save, id)
	if result.Error != nil {
		return c.Status(404).SendString("Kaydedilme Bulunamadı")
	}
	database.DB.Delete(&save)
	return c.Status(200).SendString("Kaydedilme Silindi")
}

func GetSavedPosts(c *fiber.Ctx) error {
    userID := c.Locals("userID")
    if userID == nil {
        return c.Status(fiber.StatusBadRequest).SendString("Kullanıcı ID'si bulunamadı")
    }

    userIDInt, ok := userID.(int)
    if !ok {
        return c.Status(fiber.StatusBadRequest).SendString("Geçersiz kullanıcı ID'si")
    }

    var saves []models.Save
    result := database.DB.Where("userID = ?", userIDInt).Find(&saves)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Kaydedilen yazılar alınamadı: " + result.Error.Error())
    }

   
    var postIDs []int
    for _, save := range saves {
        postIDs = append(postIDs, save.PostID)
    }

    var posts []models.Post
    result = database.DB.Where("postID IN ?", postIDs).Find(&posts)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Post detayları alınamadı: " + result.Error.Error())
    }

    return c.JSON(posts)
}

