package routers

import (
	"github.com/devgoorita/golang-basecode/internal/controllers"
	"github.com/gin-gonic/gin"
)

func (routes *Routes) addNews(
	rg *gin.RouterGroup,
	controller controllers.NewsController,
) {
	group := rg.Group("/news")
	group.GET("/paginate", controller.Pagination)
}
