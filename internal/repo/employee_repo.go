package repo

import (
	"gorm.io/gorm"
	"task/internal/models"
)

func CreateEmployeeRepo(db *gorm.DB, employee *models.Employee) error {
	if err := db.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

func CreateDesignationRepo(db *gorm.DB, designation *models.Designation) error {
	if err := db.Create(designation).Error; err != nil {
		return err
	}
	return nil
}
