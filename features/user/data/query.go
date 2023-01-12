package data

import (
	"api/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	res := User{}

	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	cekDupe := CoreToData(newUser)
	err := uq.db.Where("email=?", cekDupe.Email).First(&cekDupe).Error
	if err == nil {
		log.Println("email already registered")
		return user.Core{}, errors.New("duplicated")
	}

	cnv := CoreToData(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		return user.Core{}, err
	}

	newUser.ID = cnv.ID

	return newUser, nil
}
func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := User{}
	err := uq.db.Where("id = ?", id).First(&res).Error
	if err != nil {
		log.Println("Get By ID query error", err.Error())
		return user.Core{}, err
	}

	return ToCore(res), nil
}
func (uq *userQuery) Update(id int, updateData user.Core) (user.Core, error) {
	res := CoreToData(updateData)
	qry := uq.db.Where("id = ?", id).Updates(&res)
	if qry.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return user.Core{}, errors.New("not found")
	}
	err := qry.Error
	if err != nil {
		log.Println("update book query error :", err.Error())
		return user.Core{}, err
	}
	return ToCore(res), nil
}
func (uq *userQuery) Delete(userID int) error {

	qry := uq.db.Unscoped().Delete(&User{}, userID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no book has delete")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("book cannot delete")
	}
	return nil
}
