package repositories

import (
	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/utils"
	"github.com/devgoorita/golang-basecode/pkg"
	"gorm.io/gorm"
)

type NewsRepository struct {
	DB *gorm.DB
}

func NewNewsRepository(DB *gorm.DB) NewsRepository {
	return NewsRepository{
		DB: DB,
	}
}

func (repository *NewsRepository) List(pagination pkg.Pagination[models.News]) (*pkg.Pagination[models.News], error) {
	var news []*models.News

	repository.DB.Scopes(utils.Paginate(news, &pagination, repository.DB)).Find(&news)
	pagination.Rows = news
	return &pagination, nil
}

func (repository *NewsRepository) Create(news models.News) {
	repository.DB.Create(&news)
}

func (repository *NewsRepository) Paginate(page int, pageSize int) []models.News {
	var news []models.News
	if page <= 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	repository.DB.Offset(offset).Limit(pageSize).Find(&news)
	return news
}
