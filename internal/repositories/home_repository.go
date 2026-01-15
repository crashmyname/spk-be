package repositories

import (
	"spk-be/internal/database"
	"spk-be/internal/models"
)

func GetAllHome() ([]models.User, []models.Material, error) {
	var total int64
	var users []models.User
	var materials []models.Material

	if err := database.DB.Find(&users).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	if err := database.DB.Find(&materials).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	return users, materials, nil
}
