package services

import (
	"github.com/devgoorita/golang-basecode/internal/models"
	"github.com/devgoorita/golang-basecode/internal/repositories"
)

type RoleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(roleRepository repositories.RoleRepository) RoleService {
	return RoleService{
		roleRepository: roleRepository,
	}
}

func (service *RoleService) Create(role models.Role) {
	service.roleRepository.Create(role)
}

func (service *RoleService) FindByName(name string) (models.Role, error) {
	return service.roleRepository.FindByName(name)
}

func (service *RoleService) FindAll() []models.Role {
	return service.roleRepository.FindAll()
}
