package services

import (
	"errors"
	"services-work-team-search/database"
	"services-work-team-search/models"
)

func GetAllWorkTeam() ([]models.WorkTeam, error) {
	var data []models.WorkTeam
	result := database.DB.Where("status = ?", true).Find(&data)
	return data, result.Error
}

func GetWorkTeamByCode(teamCode string) (models.WorkTeam, error) {
	var workTeam models.WorkTeam
	result := database.DB.Where("team_code = ?", teamCode).First(&workTeam)
	if result.Error != nil {
		return workTeam, errors.New("work team not found")
	}
	return workTeam, nil
}
