package controllers

import (
	"services-work-team-search/models"
	"services-work-team-search/services"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GetWorkTeams godoc
// @Summary      List teams
// @Description  Get all the work team
// @Tags         work-team
// @Produce      json
// @Success      200  {object}  models.ApiResponse[[]models.WorkTeam]
// @Router       /work-team [get]
func GetWorkTeams(c *gin.Context) {
	list, err := services.GetAllWorkTeam()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse[[]models.WorkTeam]{
			Status:  "error",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse[[]models.WorkTeam]{
		Status:  "success",
		Data:    &list,
		Message: "List of work teams",
	})
}

// GetWorkTeamByCodigo godoc
// @Summary      Get work team by code
// @Description  Get a specific work team by its unique code
// @Tags         work-team
// @Accept       json
// @Produce      json
// @Param        team_code  path      string  true  "Unique team code"
// @Success      200            {object}  models.ApiResponse[models.WorkTeam]
// @Failure      404            {object}  models.ApiResponse[models.WorkTeam]
// @Router       /work-team/{team_code} [get]
func GetWorkTeamByCodigo(c *gin.Context) {
	teamCode := c.Param("team_code")
	workTeam, err := services.GetWorkTeamByCode(teamCode)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ApiResponse[models.WorkTeam]{
			Status:  "empty",
			Data:    &workTeam,
			Message: "Work team not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse[models.WorkTeam]{
		Status:  "success",
		Data:    &workTeam,
		Message: "Work team found",
	})
}
