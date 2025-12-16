package repositories

import (
	"errors"
	"spk-be/internal/database"
	"spk-be/internal/models"

	"gorm.io/gorm"
)

func GetAllMaterial() ([]models.Material, error) {
	var material []models.Material

	err := database.DB.Find(&material).Error

	return material, err
}

func GetMaterialByID(id int) (*models.Material, error) {
	var material models.Material

	if err := database.DB.Where("material_id = ?", id).First(&material).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &material, nil
}

func CreateMaterial(material *models.Material) error {
	return database.DB.Create(material).Error
}

func UpdateMaterial(id int, data map[string]interface{}) (*models.Material, error) {
	var material models.Material

	if err := database.DB.First(&material, id).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Model(&material).Updates(data).Error; err != nil {
		return nil, err
	}

	return &material, nil
}

func DeleteMaterial(id int) (*models.Material, error) {
	var material models.Material

	if err := database.DB.First(&material, id).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Model(&material).Delete(id).Error; err != nil {
		return nil, err
	}

	return &material, nil
}
