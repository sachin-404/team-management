package repo

import (
	"errors"
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
	team := models.TeamMember{}
	if err := db.Where("team_id = ?", teamMember.TeamID).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
	}
	if err := db.Where("team_id = ? AND user_id = ?", teamMember.TeamID, teamMember.UserID).Delete(&team).Error; err != nil {
		return err
	}
	return nil
}

func MakeAdminInDB(db *gorm.DB, userID int) error {
	user := models.User{}
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	if err := db.Model(&user).Where("id = ?", userID).Update("is_admin", true).Error; err != nil {
		return err
	}
	return nil
}
