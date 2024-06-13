package routes

import (
	"net/http"
	"task/cmd/api/handlers/auth"
	"task/cmd/api/handlers/employee"
	"task/cmd/api/handlers/team"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.POST("/login", auth.Login)
	e.POST("/signup", auth.Signup)

	e.POST("/createteam", team.CreateTeam)
	e.POST("/teams/:team_id/members/:user_id", team.AddMember)
	e.DELETE("/teams/:team_id/members/:user_id", team.RemoveMember)
	e.POST("/make_admin/:user_id", team.MakeAdmin)

	e.POST("/create_employee", employee.CreateEmployee)
	e.GET("/get_children/:id", employee.GetEmployeeHierarchy)
	e.POST("/create_designation", employee.CreateDesignation)
}
