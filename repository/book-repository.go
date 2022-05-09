package repository

import (
	"github.com/cgomez29/api-gin/models"
	"gorm.io/gorm"
)

//BookRepository is a ....
type BookRepository interface {
	InsertBook(b models.Book) models.Book
	UpdateBook(b models.Book) models.Book
	DeleteBook(b models.Book)
	AllBook() []models.Book
	FindBookByID(bookID uint64) models.Book
}

type bookConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b models.Book) models.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) UpdateBook(b models.Book) models.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b models.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) FindBookByID(bookID uint64) models.Book {
	var book models.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []models.Book {
	var books []models.Book
	db.connection.Preload("User").Find(&books)
	return books
}
