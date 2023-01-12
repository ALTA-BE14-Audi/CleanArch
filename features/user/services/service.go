package services

import (
	"api/features/user"
	"api/helper"
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password tidak sesuai")
	}

	// claims := jwt.MapClaims{}
	// claims["authorized"] = true
	// claims["userID"] = res.ID
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// useToken, _ := token.SignedString([]byte(config.JWT_KEY))
	useToken, _ := helper.GenerateToken(int(res.ID))

	return useToken, res, nil

}
func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	hashInpPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashInpPassword)
	res, err := uuc.qry.Register(newUser)
	// log.Println(res,err)
	if err != nil {
		if strings.Contains(err.Error(), "duplicated") {
			return user.Core{}, errors.New("data already exist error")
		}
		return user.Core{}, errors.New("internal server error")
	}
	log.Println("OK")
	return res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server internal error"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}
func (uuc *userUseCase) Update(token interface{}, updateData user.Core) (user.Core, error) {
	id := helper.ExtractToken(token)
	res, err := uuc.qry.Update(id, updateData)
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("query error, update fail")
	}
	return res, nil
}
func (uuc *userUseCase) Delete(token interface{}) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("data not found")
	}
	err := uuc.qry.Delete(userID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}
