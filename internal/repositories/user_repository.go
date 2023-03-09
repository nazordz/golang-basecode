package repositories

import (
	"errors"

	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{
		DB: DB,
	}
}

func (repository *UserRepository) FindAll() []models.User {
	var users []models.User
	repository.DB.Find(&users)
	return users
}

func (repository *UserRepository) Create(user models.User) (models.User, error) {
	tx := repository.DB.Create(&user)
	logrus.Info(user)
	return user, tx.Error
}

func (repository *UserRepository) FindById(id string) models.User {
	var user models.User
	repository.DB.First(&user, "id = ?", id)
	return user
}
func (repository *UserRepository) FindByColumn(column string, id string) models.User {
	var user models.User
	repository.DB.First(&user, column+" = ?", id)
	return user
}

func (repository *UserRepository) Authentication(email string, password string) (models.User, error) {
	var user models.User
	tx := repository.DB.Preload("Role").First(&user, "email = ?", email)

	if tx.Error != nil {
		return user, errors.New("email not found")
	}

	err := utils.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return user, errors.New("wrong password")
	}
	return user, nil
}
