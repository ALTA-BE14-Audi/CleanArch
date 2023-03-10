package handler

import (
	"api/features/user"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil login", res, token))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format input salah")
		}

		res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			}
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusCreated, "berhasil daftar", res))
	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", res))
	}
}
func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := uc.srv.Update(c.Get("user"), *ToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusCreated, "update successdul", res))

	}
}
func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uc.srv.Delete(c.Get("user"))
		if err != nil {
			log.Println("fail to delete")
			if strings.Contains(err.Error(), "fail") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "fail to delete, account id not found",
				})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "server error",
				})
			}

		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "deleting account successful",
		})
	}
}
