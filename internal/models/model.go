package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type Team struct {
	gorm.Model
	Name string `gorm:"uniqueIndex" json:"name"`
}

type TeamMember struct {
	gorm.Model
	TeamID uint `json:"team_id"`
	UserID uint `json:"user_id"`
}

type Employee struct {
	gorm.Model
	EmployeeID       uint        `gorm:"uniqueIndex" json:"employee_id"`
	EmployeeName     string      `json:"employee_name"`
	ParentEmployeeID *uint       `json:"parent_employee_id"`
	DesignationID    uint        `json:"designation_id"`
	Designation      Designation `gorm:"foreignKey:DesignationID;references:DesignationId" json:"-"`
}

type Designation struct {
	gorm.Model
	DesignationId uint   `gorm:"uniqueIndex" json:"designation_id"`
	Designation   string `json:"designation"`
}
