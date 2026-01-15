package services

import (
	"spk-be/internal/models"
	"spk-be/internal/repositories"
)

func GetAllDetailTicket() ([]models.DetailTicket, error) {
	return repositories.GetAllDetailTickets()
}

func GetDetailByID(id int) (models.DetailTicket, error) {
	return repositories.GetDetailByID(id)
}

func CreateDetailTicket(detailticket *models.DetailTicket) error {
	return repositories.CreateDetailTicket(detailticket)
}

func UpdateDetailTicket(id int, data map[string]interface{}) (*models.DetailTicket, error) {
	return repositories.UpdateDetailTicket(id, data)
}

func DeleteDetailTicket(id int) (*models.DetailTicket, error) {
	return repositories.DeleteDetailTicket(id)
}
