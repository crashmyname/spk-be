package repositories

import (
	"spk-be/internal/database"
	"spk-be/internal/models"
)

func GetAllTicket() ([]models.Ticket, error) {
	var tickets []models.Ticket

	if err := database.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil
}

func GetTicketByID(id int) (models.Ticket, error) {
	var tickets models.Ticket

	if err := database.DB.First(&tickets).Where("ticket_id = ?", id).Error; err != nil {
		return tickets, err
	}

	return tickets, nil
}

func CreateTicket(ticket *models.Ticket) error {
	return database.DB.Create(&ticket).Error
}

func UpdateTicket(id int, data map[string]interface{}) (*models.Ticket, error) {
	var ticket models.Ticket

	if err := database.DB.First(&ticket).Where("ticket_id = ?", id).Updates(data).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}

func DeleteTicket(id int) (*models.Ticket, error) {
	var ticket models.Ticket

	if err := database.DB.First(&ticket).Where("ticket_id = ?", id).Delete(&ticket).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}
