package main

import (
	"hello-world/backend/database"
	"hello-world/backend/handlers"
	"hello-world/backend/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Fatal("Hata: .env dosyası yüklenemedi", err)
    }

    app := fiber.New()


    database.Connect()


    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173", 
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    }))

   
    app.Get("/categories", handlers.GetCategory)
    app.Post("/category", middleware.Protected(), handlers.CreateCategory)
    app.Put("/category/:id", middleware.Protected(), handlers.UpdateCategory)
    app.Delete("/category/:id", middleware.Protected(), handlers.DeleteCategory)

    
    app.Get("/posts", handlers.GetPost)
    app.Post("/posts", middleware.Protected(), handlers.CreatePost)
    app.Put("/posts/:id", middleware.Protected(), handlers.UpdatePost)
    app.Delete("/posts/:id", middleware.Protected(), handlers.DeletePost)

    app.Get("/tags", handlers.GetTags)
    app.Post("/tags", middleware.Protected(), handlers.CreateTag)
    app.Put("/tags/:id", middleware.Protected(), handlers.UpdateTag)
    app.Delete("/tags/:id", middleware.Protected(), handlers.DeleteTag)

    app.Get("/comments", handlers.GetComments)
    app.Post("/comments", middleware.Protected(), handlers.CreateComment)
    app.Put("/comments/:id", middleware.Protected(), handlers.UpdateComment)
    app.Delete("/comments/:id", middleware.Protected(), handlers.DeleteComment)
    app.Get("/comments/:id", handlers.GetCommentById)
    

  
    app.Post("/login", handlers.LoginHandler)
    app.Post("/register", handlers.RegisterHandler)
    app.Get("/user/posts", middleware.Protected(), handlers.GetUserPosts)
    app.Get("/user/saved", middleware.Protected(), handlers.GetSavedPosts)


    app.Get("/posts/:id",handlers.GetPostById)
    app.Post("/api/save", handlers.CreateSave)



    log.Fatal(app.Listen(":3001"))
}
