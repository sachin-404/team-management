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
