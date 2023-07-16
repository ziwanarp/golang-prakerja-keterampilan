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
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: film,
	})
}

func DetailFilmController(c echo.Context) error {
	// var id, _ = strconv.Atoi(c.Param("id"))
	// var film []models.Film

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

func FilmController(c echo.Context) error {
	var films []models.Film 
	
	result := config.DB.Find(&films)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: films,
	})
}