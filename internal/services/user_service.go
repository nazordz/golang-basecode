package services

import (
	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) FindAll() []models.User {
	return userService.userRepository.FindAll()
}

func (userService *UserService) Create(user models.User) (models.User, error) {
	hasPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hasPassword)
	return userService.userRepository.Create(user)
}

func (userService *UserService) FindById(id string) models.User {
	return userService.userRepository.FindById(id)
}

func (userService *UserService) Authentication(email string, password string) (models.User, error) {
	return userService.userRepository.Authentication(email, password)

}
