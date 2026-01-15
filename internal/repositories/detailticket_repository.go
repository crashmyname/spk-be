package repositories

import (
	"spk-be/internal/database"
	"spk-be/internal/models"
)

func GetAllDetailTickets() ([]models.DetailTicket, error) {
	var detailtickets []models.DetailTicket

	if err := database.DB.Find(&detailtickets).Error; err != nil {
		return nil, err
	}

	return detailtickets, nil
}

func GetDetailByID(id int) (models.DetailTicket, error) {
	var detailtickets models.DetailTicket

	if err := database.DB.First(&detailtickets).Where("detail_id = ?", id).Error; err != nil {
		return detailtickets, err
	}

	return detailtickets, nil
}

func CreateDetailTicket(detailticket *models.DetailTicket) error {
	return database.DB.Create(&detailticket).Error
}

func UpdateDetailTicket(id int, data map[string]interface{}) (*models.DetailTicket, error) {
	var detailticket models.DetailTicket

	if err := database.DB.First(&detailticket).Where("detail_id = ?", id).Updates(data).Error; err != nil {
		return nil, err
	}

	return &detailticket, nil
}

func DeleteDetailTicket(id int) (*models.DetailTicket, error) {
	var detailticket models.DetailTicket

	if err := database.DB.First(&detailticket).Where("detail_id = ?", id).Delete(&detailticket).Error; err != nil {
		return nil, err
	}

	return &detailticket, nil
}
