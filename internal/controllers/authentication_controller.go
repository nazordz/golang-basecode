package controllers

import (
	"net/http"
	"strconv"

	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/services"
	"github.com/devgoorita/golang-basecode/internal/utils"
	"github.com/devgoorita/golang-basecode/pkg"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	userService services.UserService
}

func NewAuthenticationController(userService services.UserService) AuthenticationController {
	return AuthenticationController{
		userService: userService,
	}
}

func (controller *AuthenticationController) Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"validation": err.Error(),
		})
		return
	}

	user, err := controller.userService.Authentication(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	bearer, _ := utils.GenerateToken(user)
	refresh_token, _ := utils.GenerateRefreshToken(user.ID)
	lifetime, _ := strconv.Atoi(pkg.GodotEnv("TOKEN_HOUR_LIFESPAN"))

	c.JSON(http.StatusOK, gin.H{
		"bearer_token":  bearer,
		"refresh_token": refresh_token,
		"token_type":    "Bearer",
		"expires_in":    lifetime * 60 * 60,
	})
}

func (controller *AuthenticationController) Create(c *gin.Context) {
	var userRequest models.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqUser := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Phone:    userRequest.Phone,
		Password: userRequest.Password,
	}

	user, err := controller.userService.Create(reqUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email has been taken"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *AuthenticationController) CurrentUser(c *gin.Context) {
	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := controller.userService.FindById(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}
