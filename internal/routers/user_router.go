package routers

import (
	"github.com/devgoorita/golang-basecode/internal/controllers"
	"github.com/devgoorita/golang-basecode/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func (route *Routes) addUser(
	rg *gin.RouterGroup,
	userController controllers.UserController,
) {
	group := rg.Group("user").Use(middlewares.JwtAuthMiddleware(), middlewares.LoggedAs("admin"))
	group.GET("/", userController.FindAll)
	group.POST("/create", userController.Create)
	group.GET("/:id", userController.FindById)
}
