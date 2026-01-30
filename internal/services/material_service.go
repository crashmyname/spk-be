package services

import (
	"errors"
	"spk-be/internal/database"
	"spk-be/internal/models"
	"spk-be/internal/repositories"

	"gorm.io/gorm"
)

func GetAllMaterial() ([]models.Material, error) {
	return repositories.GetAllMaterial()
}

func GetMaterialPaginate(limit int, offset int, search string) ([]models.Material, int64, error) {
	return repositories.GetPaginateMaterial(limit, offset, search)
}

func GetMaterialByID(id string) (*models.Material, error) {
	return repositories.GetMaterialByID(id)
}

func IsMoldNumberExsist(moldNumber string) bool {
	var material models.Material
	err := database.DB.Where("mold_number= ?", moldNumber).First(&material).Error

	if err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}

	return true
}

func ImportMaterial(materials []models.Material) error {
	for _, material := range materials {
		err := repositories.CreateMaterial(&material)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateMaterial(moldnumber, lampname, modelname, typemodel string) (*models.Material, error) {
	material := &models.Material{
		MoldNumber: moldnumber,
		LampName:   lampname,
		ModelName:  modelname,
		Type:       typemodel,
	}

	err := repositories.CreateMaterial(material)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func UpdateMaterial(id string, data map[string]interface{}) (*models.Material, error) {
	return repositories.UpdateMaterial(id, data)
}

func DeleteMaterial(id int) (*models.Material, error) {
	return repositories.DeleteMaterial(id)
}
