package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kweku-xvi/todo/api/v1/dto"
	"github.com/kweku-xvi/todo/api/v1/models"
	"github.com/kweku-xvi/todo/internal/config"
	"github.com/kweku-xvi/todo/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body dto.SignUpRequest

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userFound models.User
	database.DB.Where("username = ?", body.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(400, gin.H{
			"error": "username already in use",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Name:     body.Name,
		Username: body.Username,
		Password: string(passwordHash),
	}

	database.DB.Create(&user)

	c.JSON(201, gin.H{
		"data": user,
	})
}

func SignIn(c *gin.Context) {
	var body dto.SignInRequest

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userFound models.User
	database.DB.Where("username = ?", body.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid password",
		})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(config.ENV.JWTSecret))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetUserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}
