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

func GetPaginateMaterial(limit int, offset int, search string) ([]models.Material, int64, error) {
	var materials []models.Material
	var total int64

	query := database.DB.Model(&models.Material{})

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("mold_number LIKE ? OR lamp_name LIKE ? OR model_name LIKE ? OR type LIKE ?", like, like, like, like)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("material_id DESC").Limit(limit).Offset(offset).Find(&materials).Error; err != nil {
		return nil, 0, err
	}

	return materials, total, nil
}

func GetMaterialByID(id string) (*models.Material, error) {
	var material models.Material

	if err := database.DB.Where("mold_number= ?", id).First(&material).Error; err != nil {
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

func UpdateMaterial(id string, data map[string]interface{}) (*models.Material, error) {
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

	if err := database.DB.Delete(&material).Error; err != nil {
		return nil, err
	}

	return &material, nil
}
