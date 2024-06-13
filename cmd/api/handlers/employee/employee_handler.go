package employee

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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

func GetEmployeeHierarchy(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	employeeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "employee id is invalid",
		})
	}

	rows, err := db.Raw(`
        WITH RECURSIVE employee_hierarchy AS (
            SELECT employee_id, employee_name, parent_employee_id, designation_id
            FROM employees
            WHERE employee_id = ?

            UNION ALL

            SELECT e.employee_id, e.employee_name, e.parent_employee_id, e.designation_id
            FROM employees e, employee_hierarchy
            WHERE e.parent_employee_id = employee_hierarchy.employee_id
        )

        SELECT employee_id, employee_name, parent_employee_id, designation_id FROM employee_hierarchy;
    `, employeeID).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "error getting employee hierarchy",
		})
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}(rows)

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := db.ScanRows(rows, &employee); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "error scanning row",
			})
		}
		employees = append(employees, employee)
	}

	if len(employees) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "no employees found",
		})
	}

	return c.JSON(http.StatusOK, employees)
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
