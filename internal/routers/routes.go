package routers

import (
	"github.com/devgoorita/golang-basecode/internal/controllers"
	"github.com/devgoorita/golang-basecode/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router                   *gin.Engine
	userController           controllers.UserController
	authenticationController controllers.AuthenticationController
}

func NewRoutes(
	userController controllers.UserController,
	authenticationController controllers.AuthenticationController,
) *Routes {
	r := Routes{
		router:                   gin.Default(),
		userController:           userController,
		authenticationController: authenticationController,
	}

	gin.SetMode(middlewares.GinMode())
	r.router.Use(middlewares.Gzipping)
	r.router.Use(middlewares.CorsConfig)

	api := r.router.Group("/api")
	r.addUser(api, userController)
	r.addAuthentication(api, authenticationController)

	return &r
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run()
}
