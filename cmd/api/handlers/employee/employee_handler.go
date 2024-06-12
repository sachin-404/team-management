package employee

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"task/internal/models"
	"task/internal/service"
)

func CreateEmployee(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	employee := new(models.Employee)
	if err := c.Bind(employee); err != nil {
		return err
	}

	if err := service.CreateEmployeeService(db, employee); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error creating employee",
		})
	}
	return c.JSON(http.StatusOK, employee)
}

func CreateDesignation(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	designation := new(models.Designation)
	if err := c.Bind(designation); err != nil {
		return err
	}

	if err := service.CreateDesignationService(db, designation); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error creating designation",
		})
	}
	return c.JSON(http.StatusOK, designation)
}
