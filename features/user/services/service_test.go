package services

import (
	"api/features/user"
	"api/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	pass string
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	inputData := user.Core{Nama: "audiz", Email: "audiz@mail.com", Alamat: "bangil", HP: "0814374234", Password: "asdf"}
	resData := user.Core{ID: uint(1), Nama: "audiz", Email: "audiz@mail.com", Alamat: "bangil", HP: "0814374234", Password: "asdf"}
	data.On("Register", mock.Anything).Return(resData, nil).Once()
	srv := New(data)
	res, err := srv.Register(inputData)
	assert.Nil(t, err)
	assert.Equal(t, resData.ID, res.ID)
	assert.Equal(t, resData.Nama, res.Nama)
	data.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	inputData := user.Core{Email: "audiz@mail.com", Password: "asdf"}
	resData := user.Core{ID: uint(1), Nama: "audiz", Email: "audiz@mail.com", Alamat: "bangil", HP: "0814374234", Password: "asdf"}
	data.On("Login", mock.Anything).Return(resData, nil).Once()

	// Create instance of service object and call login function
	srv := New(data)
	res, err, _ := srv.Login(inputData.Email, inputData.Password)

	// Assert that the returned result and error are as expected
	assert.Nil(t, err)
	assert.Equal(t, resData.ID, res.ID)
	assert.Equal(t, resData.Nama, res.Nama)
}

func TestProfile(t *testing.T) {
	// Set up mock object
	data := mocks.NewUserData(t)
	userID := uint(1)
	resData := user.Core{ID: userID, Nama: "audiz", Email: "audiz@mail.com", Alamat: "bangil", HP: "0814374234", Password: "asdf"}
	data.On("GetByID", userID).Return(resData, nil).Once()

	// Create instance of service object and call view profile function
	srv := New(data)
	actualResult, err := srv.Profile(userID)

	// Assert that the returned result and error are as expected
	assert.Nil(t, err)
	assert.Equal(t, resData, actualResult)
}

type mockUserData struct{}
