package controllers

import (
	"net/http"
	"golang/config"
	"golang/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: user,
	})
}

func DetailUserController(c echo.Context) error {
	// var id, _ = strconv.Atoi(c.Param("id"))
	// var user []models.User

	// query := "SELECT id, name FROM data WHERE id = ?"
	// row := db.QueryRow(query, id)

	// // Create a DataItem instance to store the retrieved data
	// item := DataItem{}
	// err = row.Scan(&item.ID, &item.Name)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return c.String(http.StatusNotFound, "Item not found")
	// 	}
	// 	log.Fatal(err)
	// }

	// return c.JSON(http.StatusOK, item)


	// var users  
	
	// result := config.DB.Find(&users)
	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, models.BaseResponse{
	// 		Message: "Gagal mendapatkan data", Status: false, Data: nil,
	// 	})
	// }
}

func UserController(c echo.Context) error {
	var users []models.User 
	
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: users,
	})
}