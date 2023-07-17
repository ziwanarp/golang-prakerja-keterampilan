package controllers

import (
	"net/http"
	"golang/config"
	"golang/models"

	"github.com/labstack/echo/v4"
)


func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error Insert Data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success Insert Data", Status: true, Data: user,
	})
}


func UserController(c echo.Context) error {
	var users []models.User 
	
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error Get Data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success Get Data", Status: true, Data: users,
	})
}