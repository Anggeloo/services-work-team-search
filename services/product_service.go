package services

import (
	"fmt"
	"time"

	"services-work-team-add/database"
	"services-work-team-add/models"
)

// generateWorkTeamCode generates a unique TeamCode based on the total count of records
func generateWorkTeamCode() (string, error) {
	var count int64
	if err := database.DB.Model(&models.WorkTeam{}).Count(&count).Error; err != nil {
		return "", err
	}
	return fmt.Sprintf("WOR_%d", count+1), nil
}

// CreateWorkTeamService creates a new work team and saves it to PostgreSQL
func CreateWorkTeamService(workTeam *models.WorkTeam) (*models.WorkTeam, error) {
	// Generate TeamCode
	teamCode, err := generateWorkTeamCode()
	if err != nil {
		return nil, err
	}
	workTeam.TeamCode = teamCode

	// Set timestamps
	workTeam.CreatedAt = time.Now()
	workTeam.UpdatedAt = time.Now()

	if err := database.DB.Create(workTeam).Error; err != nil {
		return nil, err
	}

	return workTeam, nil
}
