package controllers

import (
	"net/http"

	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/services"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return RoleController{
		roleService: roleService,
	}
}

func (controller *RoleController) Create(c *gin.Context) {
	var roleInput models.RoleInput

	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't create role"})
		return
	}
	var role models.Role
	role.Name = roleInput.Name
	go controller.roleService.Create(role)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (controller *RoleController) FindAll(c *gin.Context) {
	roles := controller.roleService.FindAll()
	c.JSON(http.StatusOK, roles)
}
