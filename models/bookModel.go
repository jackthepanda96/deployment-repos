package models

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"form:"title"`
	Author string `json:"author"form:"author"`
}

type BookDBModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) *BookDBModel {
	return &BookDBModel{db: db}
}

type BookModel interface {
	Insert(newBook Book) (Book, error)
	GetAll() ([]Book, error)
}

func (u *BookDBModel) Insert(newBook Book) (Book, error) {
	if err := u.db.Save(&newBook).Error; err != nil {
		log.Debug(err)
		return newBook, err
	}
	return newBook, nil
}

func (u *BookDBModel) GetAll() ([]Book, error) {
	book := []Book{}
	if err := u.db.Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
