package handler

import (
	"api/features/book"
	"api/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookHandle struct {
	srv book.BookService
}

func New(bs book.BookService) *bookHandle {
	return &bookHandle{
		srv: bs,
	}
}

func (bh *bookHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := bh.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan buku", res))
	}
}

func (bh *bookHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamBookID := c.Param("id")
		bookID, _ := strconv.Atoi(ParamBookID)
		input := UpdateBookReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := bh.srv.Update(c.Get("user"), bookID, *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses update buku", res))
	}
}

func (bh *bookHandle) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.srv.GetAll()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []BookResponse{}
		for i := 0; i < len(books); i++ {
			result = append(result, BookResponse(books[i]))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "", books))
	}
}

func (bh *bookHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamBookID := c.Param("id")
		bookID, _ := strconv.Atoi(ParamBookID)
		err := bh.srv.Delete(c.Get("user"), bookID)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "delete fail",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Delete successfull",
		})
	}
}

func (bh *bookHandle) MyBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("user")

		res, err := bh.srv.MyBook(userID)
		if err != nil {
			log.Println("no book found ", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "no result",
			})
		}
		result := []AddBookReponse{}
		for i := 0; i < len(res); i++ {
			result = append(result, MyBookResponse(res[i]))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "show all book list succesfull",
		})
	}
}
