package repositories

import (
	"github.com/devgoorita/golang-basecode/internal/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(DB *gorm.DB) RoleRepository {
	return RoleRepository{
		DB: DB,
	}
}

func (repository *RoleRepository) Create(role models.Role) {
	repository.DB.Create(&role)
}

func (repository *RoleRepository) FindByName(name string) (models.Role, error) {
	var role models.Role
	tx := repository.DB.First(&role, "name = ?", name)
	return role, tx.Error
}

func (repository *RoleRepository) FindAll() []models.Role {
	var roles []models.Role
	repository.DB.Find(&roles)
	return roles
}
