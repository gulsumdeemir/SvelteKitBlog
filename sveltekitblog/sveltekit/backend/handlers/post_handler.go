package handlers

import (
	"hello-world/backend/database"
	"hello-world/backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func GetPost(c *fiber.Ctx) error {
    categoryID := c.Query("categoryID")
    var posts []models.Post

    var result *gorm.DB
    if categoryID != "" {
        result = database.DB.Where("categoryID = ?", categoryID).Find(&posts)
    } else {
        result = database.DB.Find(&posts)
    }

    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Yazılar alınamadı: " + result.Error.Error())
    }

    return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
    var postInput struct {
        models.Post
        Tags []int `json:"tags"`
    }

    if err := c.BodyParser(&postInput); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz İstek"})
    }

    loggedInUser := c.Locals("userID")
    if loggedInUser == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    userID, ok := loggedInUser.(int)
    if !ok {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz kullanıcı ID'si"})
    }

    postInput.Post.UserID = userID
    postInput.Post.Datee = time.Now()

    if err := database.DB.Create(&postInput.Post).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Yazı oluşturulamadı"})
    }

    for _, tagID := range postInput.Tags {
        postTag := models.PostTag{PostID: postInput.Post.PostID, TagID: tagID}
        if err := database.DB.Create(&postTag).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostTag oluşturulamadı"})
        }
    }

    return c.Status(fiber.StatusCreated).JSON(postInput.Post) 
}

func UpdatePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var postInput struct {
        models.Post
        Tags []int `json:"tags"`
    }

 
    result := database.DB.First(&postInput.Post, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).SendString("Yazı bulunamadı")
    }

    loggedInUser := c.Locals("userID")
    if loggedInUser == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    userID, ok := loggedInUser.(int)
    if !ok || postInput.Post.UserID != userID {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    if err := c.BodyParser(&postInput); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    database.DB.Save(&postInput.Post)

 
    database.DB.Where("postID = ?", postInput.Post.PostID).Delete(&models.PostTag{})

 
    for _, tagID := range postInput.Tags {
        postTag := models.PostTag{PostID: postInput.Post.PostID, TagID: tagID}
        if err := database.DB.Create(&postTag).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "PostTag oluşturulamadı"})
        }
    }

    return c.Status(fiber.StatusOK).JSON(postInput.Post)
}

func DeletePost(c *fiber.Ctx) error {
    id := c.Params("id")
    var post models.Post
    result := database.DB.First(&post, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).SendString("Yazı bulunamadı")
    }

    loggedInUser := c.Locals("userID")
    if loggedInUser == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    userID, ok := loggedInUser.(int)
    if !ok || post.UserID != userID {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Yetkisiz"})
    }

    database.DB.Delete(&post)
    return c.Status(fiber.StatusNoContent).SendString("Yazı silindi") 
}
func GetUserPosts(c *fiber.Ctx) error {
    userID, ok := c.Locals("userID").(int)
    if !ok {
        return c.Status(fiber.StatusBadRequest).SendString("Kullanıcı ID'si bulunamadı")
    }

    var posts []models.Post
    result := database.DB.Model(&models.Post{}).Where("userID = ?", userID).Find(&posts)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Yazılar alınamadı: " + result.Error.Error())
    }
    return c.JSON(posts)
}

func GetPostById(c *fiber.Ctx) error {
    
    id := c.Params("id")
    var post models.Post
    result := database.DB.First(&post,id)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).SendString("Yazı bulunamadı")
        }

        return c.Status(fiber.StatusInternalServerError).SendString("Yazı Alınamadı: " +result.Error.Error())
    }

    return c.JSON(post)
}



