package service

import (
	"gorm.io/gorm"
	"task/internal/models"
	"task/internal/repo"
)

func CreateEmployeeService(db *gorm.DB, employee *models.Employee) error {
	if err := repo.CreateEmployeeRepo(db, employee); err != nil {
		return err
	}
	return nil
}

func CreateDesignationService(db *gorm.DB, designation *models.Designation) error {
	if err := repo.CreateDesignationRepo(db, designation); err != nil {
		return err
	}
	return nil
}
