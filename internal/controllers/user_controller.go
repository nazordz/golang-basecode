package controllers

import (
	"net/http"

	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
	roleService services.RoleService
}

func NewUserController(
	userService services.UserService,
	roleService services.RoleService,
) UserController {
	return UserController{
		userService: userService,
		roleService: roleService,
	}
}

func (controller *UserController) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, controller.userService.FindAll())
}

func (userController *UserController) Create(c *gin.Context) {
	var userRequest models.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role, err := userController.roleService.FindByName(userRequest.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role invalid"})
		return
	}
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Phone:    userRequest.Phone,
		Password: userRequest.Password,
		RoleID:   role.ID,
		Role:     role,
	}

	newUser, err := userController.userService.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email has been taken"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (userController *UserController) FindById(c *gin.Context) {
	userId := c.Param("id")
	user := userController.userService.FindById(userId)

	c.JSON(http.StatusOK, user)
}
