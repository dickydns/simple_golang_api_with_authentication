package controllers

import (
	"net/http"
	"simple_golang_api_with_authentication/config"
	"simple_golang_api_with_authentication/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(ctx *gin.Context) {
	var users []models.User

	//QUERY
	result := config.DB.Where("status =? ", 1).Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data user",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	result := config.DB.First(&user, id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User tidak ditemukan",
		})
		return
	}
	//struct data get req body
	var input models.UpdateUser

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.Nama != nil {
		user.Nama = *input.Nama
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.Password != nil {
		hashed, _ := bcrypt.GenerateFromPassword(
			[]byte(*input.Password),
			bcrypt.DefaultCost,
		)
		user.Password = string(hashed)
	}

	if input.Status != nil {
		user.Status = *input.Status
	}

	config.DB.Save(&user)
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	result := config.DB.First(&user, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User tidak ditemukan",
		})
		return
	}

	config.DB.Model(&models.User{}).Where("id =? ", id).Update("status", false)

	ctx.JSON(http.StatusOK, user)
}
