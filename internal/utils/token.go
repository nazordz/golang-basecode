package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) (string, error) {

	token_lifespan, err := strconv.Atoi(pkg.GodotEnv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(pkg.GodotEnv("JWT_SECRET_KEY")))

}

func GenerateRefreshToken(user models.User) (string, error) {
	token_lifespan, err := strconv.Atoi(pkg.GodotEnv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(pkg.GodotEnv("JWT_SECRET_KEY")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(pkg.GodotEnv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (string, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(pkg.GodotEnv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "Error ", err
	}
	claims, ok := token.Claims.(*JwtClaims)

	if ok && token.Valid {
		return claims.User.ID, nil
	}
	return "", errors.New("token invalid")
}
