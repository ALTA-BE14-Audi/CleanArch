package data

import (
	"api/features/book"
	"errors"
	"log"

	// "strconv"

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

// func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
// 	cnv := CoreToData(newBook)
// 	cnv.UserID = uint(userID)
// 	err := bd.db.Create(&cnv).Error
// 	if err != nil {
// 		return book.Core{}, err
// 	}

// 	newBook.ID = cnv.ID

// 	return newBook, nil
// }

func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bd.db.Create(&cnv).Error
	if err != nil {
		log.Println("error insert user", err.Error())
		return book.Core{}, errors.New("error insert user")
	}

	newBook.ID = cnv.ID

	//isi Pemilik pakai angka
	// IDUser := strconv.Itoa(int(cnv.UserID))
	// newBook.Pemilik = IDUser

	qry := User{}
	err2 := bd.db.Where("id=?", userID).First(&qry).Error
	if err2 != nil {
		log.Println("no id found")
		return book.Core{}, errors.New("data not found")
	}
	newBook.Pemilik = qry.Nama

	return newBook, nil
}

func (bd *bookData) Update(tokenID int, bookID int, updatedData book.Core) (book.Core, error) {
	cnv := CoreToData(updatedData)
	qry := bd.db.Where("id = ?", bookID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return book.Core{}, errors.New("not found")
	}

	err := qry.Error
	if err != nil {
		log.Println("update book query error :", err.Error())
		return book.Core{}, err
	}
	log.Println(tokenID)
	qry2 := User{}
	err2 := bd.db.Where("id=?", tokenID).First(&qry2).Error
	if err2 != nil {
		log.Println("no id found")
		return book.Core{}, errors.New("data not found")
	}
	updatedData.Pemilik = qry2.Nama
	return updatedData, nil
}
func (bd *bookData) GetAll() ([]book.Core, error) {
	var books []Books
	err := bd.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	var bookCores []book.Core
	for i := 0; i < len(books); i++ {
		temp := books[i]
		bookCores = append(bookCores, ToCore(temp))
		// cnv := strconv.Itoa(int(books[i].UserID))
		// bookCores[i].Pemilik = cnv
		qry := User{}
		err := bd.db.Where("id=?", books[i].UserID).First(&qry).Error
		if err != nil {
			log.Println("no data found")
			return []book.Core{}, errors.New("data not found")
		}
		bookCores[i].Pemilik = qry.Nama
	}

	return bookCores, nil
}

// func (bd *bookData) BookList() ([]book.Core, error) {
//   res := []Books{}
//   err := bd.db.Find(&res).Error
//   if err != nil {
//     log.Println("no data found")
//     return []book.Core{}, errors.New("data not found")
//   }
//   result := []book.Core{}
//   for i := 0; i < len(res); i++ {
//     temp := res[i]
//     result = append(result, ToCore(temp))
//     // qry := User{}
//     // err := bd.db.Where("id=?", res[i].UserID).First(&qry).Error
//     // if err != nil {
//     //   log.Println("no data found")
//     //   return []book.Core{}, errors.New("data not found")
//     // }
//     // result[i].Pemilik = qry.Name
//   }
//   // log.Println(result)
//   return result, nil
// }

func (bd *bookData) Delete(userID int, bookID int) error {
	check := []Books{}
	err := bd.db.Where("id=? AND user_id=?", bookID, userID).Find(&check).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("book not found,fail deleting")
	}
	if len(check) == 0 {
		return errors.New("book not found")
	}

	qry := bd.db.Unscoped().Delete(&Books{}, bookID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no book has delete")
	}
	err = qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("book cannot delete")
	}
	return nil
}

func (bd *bookData) MyBook(userID int) ([]book.Core, error) {
	res := []Books{}
	err := bd.db.Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		log.Println("no result")
		return []book.Core{}, errors.New("data not found")
	}
	result := []book.Core{}
	for i := 0; i < len(res); i++ {
		tmp := res[i]
		result = append(result, ToCore(tmp))
	}

	return result, nil
}

func (bd *bookData) Book(userID int) ([]book.Core, error) {
	res := []Books{}
	err := bd.db.Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		log.Println("no result")
		return []book.Core{}, errors.New("data not found")
	}
	result := []book.Core{}
	for i := 0; i < len(res); i++ {
		tmp := res[i]
		result = append(result, ToCore(tmp))
	}

	return result, nil
}
