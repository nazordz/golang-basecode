package routers

import (
	"github.com/devgoorita/golang-basecode/internal/controllers"
	"github.com/devgoorita/golang-basecode/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func (route *Routes) addRole(
	rg *gin.RouterGroup,
	roleController controllers.RoleController,

) {
	group := rg.Group("/role").Use(middlewares.JwtAuthMiddleware(), middlewares.LoggedAs("admin"))
	group.POST("/", roleController.Create)
	group.GET("/", roleController.FindAll)
}
