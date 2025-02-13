package routes

import (
	"services-work-team-add/controllers"

	"github.com/gin-gonic/gin"
)

func WorkTeamRoutes(r *gin.RouterGroup) {
	r.POST("/work-team/add", controllers.CreateWorkTeam)
}
