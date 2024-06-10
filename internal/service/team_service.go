package service

import (
	"task/internal/models"
	"task/internal/repo"

	"gorm.io/gorm"
)

func CreateTeamService(db *gorm.DB, team *models.Team) error {
	if err := repo.CreateTeamInDB(db, team); err != nil {
		return err
	}
	return nil
}

func AddMemberService(db *gorm.DB, teamMember *models.TeamMember) error {
	if err := repo.AddMemberToDB(db, teamMember); err != nil {
		return err
	}
	return nil
}

func RemoveMemberService(db *gorm.DB, teamMember *models.TeamMember) error {
	if err := repo.RemoveMemberFromDB(db, teamMember); err != nil {
		return err
	}
	return nil
}

func MakeAdminService(db *gorm.DB, userID int) error {
	if err := repo.MakeAdminInDB(db, userID); err != nil {
		return err
	}
	return nil
}
