package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nikodemus.com/controllers"
	"nikodemus.com/models"
)

func main() {

	r := gin.Default()

	//MODEL
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Universitas AKI API"})
	})

	r.GET("/mahasiswa", controllers.MahasiswaTampil)
	r.POST("/mahasiswa", controllers.MahasiswaTambah)
	r.PUT("/mahasiswa/:nim", controllers.MahasiswaUbah)
	r.DELETE("/mahasiswa/:nim", controllers.MahasiswaHapus)

	r.Run()
}
