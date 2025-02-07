package routes

import (
	"services-work-team-search/controllers"

	"github.com/gin-gonic/gin"
)

func WorkTeamRoutes(r *gin.RouterGroup) {
	r.GET("/work-team", controllers.GetWorkTeams)
	r.GET("/work-team/:team_code", controllers.GetWorkTeamByCodigo)
}
