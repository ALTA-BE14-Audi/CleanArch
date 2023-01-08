package data

import (
	"api/features/book"
	"errors"
	"log"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bd.db.Create(&cnv).Error
	if err != nil {
		return book.Core{}, err
	}

	newBook.ID = cnv.ID

	return newBook, nil
}
func (bd *bookData) Update(bookID int, updatedData book.Core) (book.Core, error) {
	cnv := CoreToData(updatedData)
	qry := bd.db.Where("id = ?", bookID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return book.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update book query error :", err.Error())
		return book.Core{}, err
	}

	return ToCore(cnv), nil
}
func (bd *bookData) GetAll() ([]book.Core, error) {
	var books []Books
	err := bd.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	var bookCores []book.Core
	for _, b := range books {
		bookCores = append(bookCores, ToCore(b))
	}

	return bookCores, nil
}

// func (bd *bookData) Delete(bookID int, userID int) error {
// 	return nil
// }
// func (bd *bookData) MyBook(userID int) ([]book.Core, error) {
// 	return nil, nil
// }
