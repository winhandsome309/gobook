package models

import (
	"time"

	"github.com/winhandsome309/gobook/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string    `json:"name" gorm:"default:''"`
	Author      string    `json:"author" gorm:"default:''"`
	Publication time.Time `json:"publication"`
}

func Init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBook() *[]Book {
	var books []Book
	db.Find(&books)
	return &books
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBookById(bookId int) *Book {
	b := &Book{}
	db.Where("id = ?", bookId).First(b)
	return b
}

func (b *Book) UpdateBook(bNew *Book) {
	if bNew.Name != "" {
		b.Name = bNew.Name
	}
	if bNew.Author != "" {
		b.Author = bNew.Author
	}
	if !bNew.Publication.IsZero() {
		b.Publication = bNew.Publication
	}
	db.Save(&b)
}

func DeleteBook(bookId int) {
	// Delete softly
	// db.Where("id = ?", bookId).Delete(&Book{})
	// Delete permanently
	db.Where("id = ?", bookId).Unscoped().Delete(&Book{})
}
