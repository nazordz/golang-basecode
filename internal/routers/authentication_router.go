package routers

import (
	"github.com/devgoorita/golang-basecode/internal/controllers"
	"github.com/gin-gonic/gin"
)

func (route *Routes) addAuthentication(
	rg *gin.RouterGroup,
	authenticationController controllers.AuthenticationController,
) {
	group := rg.Group("authentication")
	group.POST("/login", authenticationController.Login)
	group.POST("/register", authenticationController.Create)
	group.GET("/current", authenticationController.CurrentUser)

}
