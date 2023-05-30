package controllers

import (
	"net/http"
	"strconv"

	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewsController struct {
	newsService services.NewsService
	DB          *gorm.DB
}

func NewNewsController(newsService services.NewsService, DB *gorm.DB) NewsController {
	return NewsController{
		newsService: newsService,
		DB:          DB,
	}
}

func (controller *NewsController) Pagination(c *gin.Context) {

	page, errPage := strconv.Atoi(c.Query("page"))
	pageSize, errPerPage := strconv.Atoi(c.Query("per_page"))
	// c.String(http.StatusOK, "Hello %s %s", page, pageSize)
	if errPage != nil {
		page = 1
	}
	if errPerPage != nil {
		pageSize = 2
	}

	var news []models.News
	// if page <= 0 {
	// 	page = 1
	// }
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	controller.DB.Offset(offset).Limit(pageSize).Find(&news)
	c.JSON(http.StatusOK, news)
}
