package services

import (
	"spk-be/internal/models"
	"spk-be/internal/repositories"
)

func GetAllTicket() ([]models.Ticket, error) {
	return repositories.GetAllTicket()
}

func GetTicketByID(id int) (models.Ticket, error) {
	return repositories.GetTicketByID(id)
}

func CreateTicket(ticket *models.Ticket) error {
	return repositories.CreateTicket(ticket)
}

func UpdateTicket(id int, data map[string]interface{}) (*models.Ticket, error) {
	return repositories.UpdateTicket(id, data)
}

func DeleteTicket(id int) (*models.Ticket, error) {
	return repositories.DeleteTicket(id)
}
