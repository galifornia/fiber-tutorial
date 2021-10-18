package main

import (
	"fmt"

	"github.com/galifornia/fiber-tutorial/book"
	"github.com/galifornia/fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connection is open")

	// Migrate the schema
	database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	// !FIXME: cannot close with v2 of gorm
	// defer database.DBConn.Close()

	setupRoutes(app)

	app.Get("/", hello)

	app.Listen(":8000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)

	app.Post("/api/v1/book", book.NewBook)

	app.Put("/api/v1/book/:id", book.UpdateBook)

	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func hello(c *fiber.Ctx) {
	c.SendString("Hello, World")
}
