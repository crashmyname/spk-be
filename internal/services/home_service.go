package services

import (
	"spk-be/internal/models"
	"spk-be/internal/repositories"
)

func GetAllHome() ([]models.User, []models.Material, error) {
	return repositories.GetAllHome()
}
