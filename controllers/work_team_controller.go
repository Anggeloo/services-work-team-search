package controllers

import (
	"net/http"

	"services-work-team-add/models"
	"services-work-team-add/services"

	"github.com/gin-gonic/gin"
)

// CreateWorkTeam godoc
// @Summary      Create a work team
// @Description  Create a new work team
// @Tags         work-team
// @Accept       json
// @Produce      json
// @Param        workTeam body models.WorkTeam true "Work team data"
// @Success      200  {object}  models.ApiResponse[models.WorkTeam]
// @Failure      400  {object}  models.ApiResponse[models.WorkTeam]
// @Failure      500  {object}  models.ApiResponse[models.WorkTeam]
// @Router       /work-team/add [post]
func CreateWorkTeam(c *gin.Context) {
	var workTeam models.WorkTeam

	// Bind JSON request to WorkTeam struct
	if err := c.ShouldBindJSON(&workTeam); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse[models.WorkTeam]{
			Status:  "error",
			Data:    nil,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Call the service to create the work team
	createdWorkTeam, err := services.CreateWorkTeamService(&workTeam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse[models.WorkTeam]{
			Status:  "error",
			Data:    nil,
			Message: "Error creating work team: " + err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, models.ApiResponse[models.WorkTeam]{
		Status:  "success",
		Data:    createdWorkTeam,
		Message: "Work team successfully created",
	})
}
