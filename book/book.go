package book

import (
	"github.com/galifornia/fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(&book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book has been found")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted")
}

func UpdateBook(c *fiber.Ctx) {
	c.SendString("update book")
}
