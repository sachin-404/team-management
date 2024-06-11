package team

import (
	"net/http"
	"strconv"
	"task/internal/models"
	"task/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTeam(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	team := new(models.Team)
	if err := c.Bind(team); err != nil {
		return err
	}

	if err := service.CreateTeamService(db, team); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create team",
		})
	}
	return c.JSON(http.StatusOK, team)
}

func AddMember(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid team ID",
		})
	}
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	teamMember := models.TeamMember{
		TeamID: uint(teamID),
		UserID: uint(userID),
	}

	if err := service.AddMemberService(db, &teamMember); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not add member",
		})
	}
	return c.JSON(http.StatusOK, teamMember)
}

func RemoveMember(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	teamID, err := strconv.Atoi(c.Param("team_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid team ID",
		})
	}
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	teamMember := models.TeamMember{
		TeamID: uint(teamID),
		UserID: uint(userID),
	}

	if err := service.RemoveMemberService(db, &teamMember); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not remove member",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Member removed",
	})
}

func MakeAdmin(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	if err := service.MakeAdminService(db, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not make admin",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User is now an admin",
	})
}
