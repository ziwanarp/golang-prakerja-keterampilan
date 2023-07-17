package controllers

import (
	"net/http"
	"golang/config"
	"golang/models"
	// "strconv"

	"github.com/labstack/echo/v4"
)


func CreateFilmController(c echo.Context) error {
	film := models.Film{}
	c.Bind(&film)

	result := config.DB.Create(&film)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error Insert Data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success Insert Data", Status: true, Data: film,
	})
}

func FilmController(c echo.Context) error {
	var films []models.Film 
	
	result := config.DB.Find(&films)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error Get Data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success Get Data", Status: true, Data: films,
	})
}