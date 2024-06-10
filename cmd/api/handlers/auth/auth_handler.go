package auth

import (
	"net/http"
	"task/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("lojXHYfs32g4pMF5i4J6glO1slQaZTND48x8Jycq9es=")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Signup(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create user",
		})
	}
	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	dbUser := new(models.User)
	db.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.ID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	claims := &JWTClaims{
		Username: dbUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create token",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
