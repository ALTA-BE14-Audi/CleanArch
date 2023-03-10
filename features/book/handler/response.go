package handler

import (
	"api/features/book"
	"api/features/user"
	"net/http"
	"strings"
)

type UserReponse struct {
	ID     uint   `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	HP     string `json:"hp"`
}
type BookResponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
	Pemilik     string `json:"pemilik"`
}

type BookList struct {
	Judul   string `json:"title"`
	Penulis string `json:"written by"`
	Pemilik string `json:"owner name"`
}

type UpdateBook struct {
	Judul       string `json:"title"`
	TahunTerbit int    `json:"published_year"`
	Penulis     string `json:"written by"`
}

type AddBookReponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"pemilik"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:     data.ID,
		Nama:   data.Nama,
		Email:  data.Email,
		Alamat: data.Alamat,
		HP:     data.HP,
	}
}

func MyBookResponse(data book.Core) AddBookReponse {
	return AddBookReponse{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = ToResponse(data[0].(user.Core))
	} else {
		resp["data"] = ToResponse(data[0].(user.Core))
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "not found") {
		code = http.StatusNotFound
	}

	return code, resp
}
