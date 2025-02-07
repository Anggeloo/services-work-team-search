package models

import (
	"time"

	"github.com/lib/pq"
)

// ApiResponse Standard structure for API responses
// @Description Generic API response
type ApiResponse[T any] struct {
	Status  string `json:"status" example:"success"`               // Response status
	Data    *T     `json:"data"`                                   // Response data
	Message string `json:"message" example:"Successful operation"` // Descriptive message
}

// WorkTeam Work team model
// @Description Structure representing a work team
type WorkTeam struct {
	WorkTeamId  uint           `json:"work_team_id" gorm:"primaryKey;autoIncrement" example:"1"`
	TeamCode    string         `json:"team_code" gorm:"type:varchar(100);unique;not null" example:"TEAM-001"`
	TeamName    string         `json:"team_name" gorm:"type:varchar(255);not null" example:"Development Team"`
	Description string         `json:"description" gorm:"type:text" example:"Team responsible for software development"`
	Members     pq.StringArray `json:"members" gorm:"type:text[]" swaggertype:"array,string" example:"[\"user1\", \"user2\"]"`
	Status      bool           `json:"status" gorm:"default:true" example:"true"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime" example:"2023-01-01T00:00:00Z"`
}
