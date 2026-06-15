package repository

import (
	"fmt"
	"http/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(row models.UserRoles) (models.UserRoles, error)
	Get(row []models.UserRoles) ([]models.UserRoles, error)
	GetById(row models.UserRoles, id int) (models.UserRoles, error)
	Update(row models.UserRoles, id int) (models.UserRoles, error)
	UpdateAll(row models.UserRoles, id int) (models.UserRoles, error)
	Delete(row models.UserRoles, id int) (models.UserRoles, error)
}

type database struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &database{DB: db}
}

func (d database) Create(row models.UserRoles) (models.UserRoles, error) {

	if err := d.DB.Create(&row).Error; err != nil {
		return row, fmt.Errorf("Failed inserting the row : %v", err)
	}

	return row, nil
}

func (d database) Get(row []models.UserRoles) ([]models.UserRoles, error) {

	if err := d.DB.Preload("UserRoleMapping").Find(&row).Error; err != nil {
		return row, fmt.Errorf("Failed fetching the rows : %v", err)
	}

	return row, nil
}

func (d database) GetById(row models.UserRoles, id int) (models.UserRoles, error) {

	if err := d.DB.Preload("UserRoleMapping").First(&row, id).Error; err != nil {
		return row, fmt.Errorf("ID not found : %v", err)
	}

	return row, nil
}

func (d database) Delete(row models.UserRoles, id int) (models.UserRoles, error) {

	if err := d.DB.First(&row, id).Error; err != nil {
		return row, fmt.Errorf("ID not found : %v", err)
	}

	if err := d.DB.Delete(&row, id).Error; err != nil {
		return row, fmt.Errorf("Failed deleting the row : %v", err)
	}

	return row, nil
}

func (d database) UpdateAll(row models.UserRoles, id int) (models.UserRoles, error) {

	if err := d.DB.First(&row, id).Error; err != nil {
		return row, fmt.Errorf("ID not found : %v", err)
	}

	if err := d.DB.Save(&row).Error; err != nil {
		return row, fmt.Errorf("Failed updating the row : %v", err)
	}

	return row, nil
}

func (d database) Update(row models.UserRoles, id int) (models.UserRoles, error) {

	if err := d.DB.First(&row, id).Error; err != nil {
		return row, fmt.Errorf("ID not found : %v", err)
	}

	if err := d.DB.Model(&models.UserRoles{}).Where("id =?", id).Updates(row).Error; err != nil {
		return row, fmt.Errorf("Failed updating the row : %v", err)
	}

	return row, nil
}
