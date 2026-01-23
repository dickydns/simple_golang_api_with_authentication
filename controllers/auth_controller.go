package controllers

import (
	"net/http"
	"os"
	"simple_golang_api_with_authentication/config"
	"simple_golang_api_with_authentication/models"
	"simple_golang_api_with_authentication/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "JSON TIDAK VALID",
		})
		return
	}

	var user models.User
	//CHECK EMAIL
	result := config.DB.Where("email = ? ", input.Email).Find(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau password salah",
		})
		return
	}

	//CHECK PASSWORD
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau password salah",
		})
		return
	}

	//JWT GENERATE
	token, err := utils.GenerateToken(user.ID)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal generate token",
		})
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusOK, gin.H{
		"msg":           "Login berhasil",
		"token":         token,
		"refresh_token": refreshToken,
		"user":          user,
	})

}

func RegisterUser(ctx *gin.Context) {
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//hash password
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Nama:     input.Nama,
		Email:    input.Email,
		Password: string(hashPassword),
		Status:   true,
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

func RefreshToken(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	claims := &utils.RefreshClaims{}

	token, err := jwt.ParseWithClaims(
		input.RefreshToken,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
		},
	)

	if err != nil || !token.Valid {
		ctx.JSON(401, gin.H{"error": "Refresh token invalid"})
		return
	}

	// Generate access token baru
	newAccessToken, _ := utils.GenerateToken(claims.UserID)

	ctx.JSON(200, gin.H{
		"access_token": newAccessToken,
	})
}
