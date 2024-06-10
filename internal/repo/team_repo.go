package repo

import (
	"task/internal/models"

	"gorm.io/gorm"
)

func CreateTeamInDB(db *gorm.DB, team *models.Team) error {
	if err := db.Create(team).Error; err != nil {
		return err
	}
	return nil
}

func AddMemberToDB(db *gorm.DB, teamMember *models.TeamMember) error {
	if err := db.Create(&teamMember).Error; err != nil {
		return err
	}
	return nil
}

func RemoveMemberFromDB(db *gorm.DB, teamMember *models.TeamMember) error {
	if err := db.Where("team_id = ? AND user_id = ?", teamMember.TeamID, teamMember.UserID).Delete(&models.TeamMember{}).Error; err != nil {
		return err
	}
	return nil
}

func MakeAdminInDB(db *gorm.DB, userID int) error {
	if err := db.Model(&models.User{}).Where("id = ?", userID).Update("is_admin", true).Error; err != nil {
		return err
	}
	return nil
}
