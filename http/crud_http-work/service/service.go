package service

import (
	"http/models"
	"http/repository"
)

type Service interface {
	CreateUserRoles(row models.UserRoles) (models.UserRoles, error)
	GetUserRoles(row []models.UserRoles) ([]models.UserRoles, error)
	GetByIdUserRoles(row models.UserRoles, id int) (models.UserRoles, error)
	UpdateUserRoles(row models.UserRoles, id int) (models.UserRoles, error)
	UpdateAllUserRoles(row models.UserRoles, id int) (models.UserRoles, error)
	DeleteUserRoles(row models.UserRoles, id int) (models.UserRoles, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {

	return &service{repo: repo}

}

func (s *service) CreateUserRoles(row models.UserRoles) (models.UserRoles, error) {

	result, err := s.repo.Create(row)
	if err != nil {
		return row, err
	}

	return result, nil
}

func (s *service) GetUserRoles(row []models.UserRoles) ([]models.UserRoles, error) {

	result, err := s.repo.Get(row)
	if err != nil {
		return row, err
	}

	return result, nil
}

func (s *service) GetByIdUserRoles(row models.UserRoles, id int) (models.UserRoles, error) {

	result, err := s.repo.GetById(row, id)
	if err != nil {
		return row, err
	}

	return result, nil
}

func (s *service) UpdateUserRoles(row models.UserRoles, id int) (models.UserRoles, error) {

	result, err := s.repo.GetById(row, id)
	if err != nil {
		return row, err
	}

	return result, nil
}

func (s *service) UpdateAllUserRoles(row models.UserRoles, id int) (models.UserRoles, error) {

	result, err := s.repo.GetById(row, id)
	if err != nil {
		return row, err
	}

	return result, nil
}

func (s *service) DeleteUserRoles(row models.UserRoles, id int) (models.UserRoles, error) {

	result, err := s.repo.GetById(row, id)
	if err != nil {
		return row, err
	}

	return result, nil
}
