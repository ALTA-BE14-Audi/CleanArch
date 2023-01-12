package services

import (
	"api/features/user"
	"api/helper"
	"api/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	inputData := user.Core{ID: uint(0), Nama: "audiz", Email: "audiz@mail.com", Alamat: "bangil", HP: "0814374234", Password: "asdf"}
	resData := user.Core{ID: uint(1), Nama: "audiz", Email: "audiz@mail.com", Alamat: "bangil", HP: "0814374234", Password: "asdf"}
	t.Run("Berhasil register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Nama, res.Nama)
		data.AssertExpectations(t)
	})
	t.Run("Gagal register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("error saat daftar"))
		srv := New(data)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.NotEqual(t, res.Nama, resData.Nama)
		data.AssertExpectations(t)
	})
	// t.Run("Duplikat", func(t *testing.T) {
	// 	data.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()
	// 	srv := New(data)
	// 	res, err := srv.Register(user.Core{})
	// 	assert.ErrorContains(t, err, "error")
	// 	assert.Equal(t, "", res.Password)
	// 	data.AssertExpectations(t)
	// })
}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	inputEmail := "audiz@gmail.com"
	hashed, _ := helper.Generate("123")
	// res dari data akan mengembalikan password yang sudah di hash
	resData := user.Core{ID: uint(1), Nama: "audi", Email: "audiz@gmail.com", HP: "08123456", Password: hashed}
	t.Run("Berhasil login", func(t *testing.T) {
		data.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(data)
		token, res, err := srv.Login(inputEmail, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Tidak ditemukan", func(t *testing.T) {
		data.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(data)
		token, res, err := srv.Login(inputEmail, "123")
		assert.NotNil(t, err)
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Salah password", func(t *testing.T) {
		inputEmail := "audiz@gmail.com"
		hashed, _ := helper.Generate("be1422")
		resData := user.Core{ID: uint(1), Nama: "jerry", Email: "jerry@alterra.id", HP: "08123456", Password: hashed}
		data.On("Login", inputEmail).Return(resData, nil).Once()

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1423")
		assert.NotNil(t, err)
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Nama: "joe", Email: "joe@gmail.com", HP: "0147234"}
		data.On("Profile", uint(1)).Return(resData, nil).Once()
		srv := New(data)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})
	t.Run("JWT tidak valid", func(t *testing.T) {
		srv := New(data)
		_, token := helper.GenerateToken(1)
		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("Data tidak ditemukan", func(t *testing.T) {
		data.On("Profile", uint(4)).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(data)

		_, token := helper.GenerateToken(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Server error", func(t *testing.T) {
		data.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(data)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
func TestUpdate(t *testing.T) {
	data := mocks.NewUserData(t)
	inputData := user.Core{ID: uint(0), Nama: "audiz", Email: "audiz@gmail.com", Alamat: "bangil", HP: "0922388923"}
	resData := user.Core{
		ID: uint(1), Nama: "audiz", Email: "audiz@gmail.com", Alamat: "bangil", HP: "0922388923",
	}
	t.Run("Sukses update", func(t *testing.T) {
		data.On("Update", int(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		token := tokenIDUser.(*jwt.Token)
		token.Valid = true
		res, err := srv.Update(token, inputData)
		assert.Equal(t, resData.ID, res.ID)
		assert.Nil(t, err)
		data.AssertExpectations(t)

	})
	t.Run("Gagal update", func(t *testing.T) {
		data.On("Update", int(1), inputData).Return(resData, errors.New("query error")).Once()
		srv := New(data)
		_, tokenID := helper.GenerateToken(1)
		token := tokenID.(*jwt.Token)
		token.Valid = true
		res, err := srv.Update(token, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "fail")
		assert.Equal(t, "", res.Nama)
		data.AssertExpectations(t)
	})
}
func TestDelete(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("Sukses hapus", func(t *testing.T) {
		data.On("Delete", int(1)).Return(nil).Once()
		_, token := helper.GenerateToken(1)
		IDToken := token.(*jwt.Token)
		IDToken.Valid = true
		srv := New(data)
		err := srv.Delete(IDToken)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})
	t.Run("Hapus gagal", func(t *testing.T) {
		data.On("Delete", int(1)).Return(errors.New("fail to delete")).Once()
		_, token := helper.GenerateToken(1)
		IDToken := token.(*jwt.Token)
		IDToken.Valid = true
		srv := New(data)
		err := srv.Delete(IDToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "fail")
		data.AssertExpectations(t)
	})
	t.Run("ID salah", func(t *testing.T) {
		_, token := helper.GenerateToken(1)
		srv := New(data)
		err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

}
