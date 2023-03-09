//go:build wireinject
// +build wireinject

package main

import (
	"github.com/devgoorita/golang-basecode/internal/controllers"
	"github.com/devgoorita/golang-basecode/internal/repositories"
	"github.com/devgoorita/golang-basecode/internal/routers"
	"github.com/devgoorita/golang-basecode/internal/services"
	"github.com/devgoorita/golang-basecode/pkg"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repositories.NewUserRepository,
	services.NewUserService,
	controllers.NewUserController,
)

var roleSet = wire.NewSet(
	repositories.NewRoleRepository,
	services.NewRoleService,
	controllers.NewRoleController,
)

func InitializedServer() *routers.Routes {
	wire.Build(
		pkg.NewDB,
		userSet,
		roleSet,
		controllers.NewAuthenticationController,
		routers.NewRoutes,
	)

	return nil
}
