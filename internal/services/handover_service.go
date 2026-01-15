package services

import (
	"spk-be/internal/models"
	"spk-be/internal/repositories"
)

func GetAllHandover() ([]models.Handover, error) {
	return repositories.GetAllHandover()
}

func GetHandoverByID(id int) (models.Handover, error) {
	return repositories.GetHandoverByID(id)
}

func CreateHandover(handover *models.Handover) error {
	return repositories.CreateHandover(handover)
}

func UpdateHandover(id int, data map[string]interface{}) (*models.Handover, error) {
	return repositories.UpdateHandover(id, data)
}

func DeleteHandover(id int) (*models.Handover, error) {
	return repositories.DeleteHandover(id)
}
