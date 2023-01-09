package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Judul       string
	TahunTerbit int
	Penulis     string
	Pemilik     string
}

type BookHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	MyBook() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	Update(token interface{}, bookID int, updatedData Core) (Core, error)
	GetAll() ([]Core, error)
	Delete(token interface{}, bookID int) error
	MyBook(token interface{}) ([]Core, error)
}

type BookData interface {
	Add(userID int, newBook Core) (Core, error)
	Update(token int, bookID int, updatedData Core) (Core, error)
	GetAll() ([]Core, error)
	Delete(bookID int, userID int) error
	MyBook(userID int) ([]Core, error)
}
