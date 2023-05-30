package services

import (
	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/repositories"
)

type NewsService struct {
	newsRepository repositories.NewsRepository
}

func NewNewsService(newsRepository repositories.NewsRepository) NewsService {
	return NewsService{
		newsRepository: newsRepository,
	}
}

func (service *NewsService) Paginate(page int, size int) []models.News {
	return service.newsRepository.Paginate(page, size)
}

func (service *NewsService) Create(news models.News) {
	// service.newsRepository.Create()
}
