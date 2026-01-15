package repositories

import (
	"spk-be/internal/database"
	"spk-be/internal/models"
)

func GetAllHandover() ([]models.Handover, error) {
	var handovers []models.Handover

	if err := database.DB.Find(&handovers).Error; err != nil {
		return nil, err
	}

	return handovers, nil
}

func GetHandoverByID(id int) (models.Handover, error) {
	var handover models.Handover

	if err := database.DB.First(&handover).Where("handover_id = ?", id).Error; err != nil {
		return handover, err
	}

	return handover, nil
}

func CreateHandover(handover *models.Handover) error {
	return database.DB.Create(&handover).Error
}

func UpdateHandover(id int, data map[string]interface{}) (*models.Handover, error) {
	var handover models.Handover

	if err := database.DB.First(&handover).Where("handover_id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	return &handover, nil
}

func DeleteHandover(id int) (*models.Handover, error) {
	var handover models.Handover

	if err := database.DB.First(&handover).Where("handover_id = ?", id).Delete(&handover).Error; err != nil {
		return nil, err
	}

	return &handover, nil
}
