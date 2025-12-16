package services

import (
	"spk-be/internal/models"
	"spk-be/internal/repositories"
)

func GetAllMaterial() ([]models.Material, error) {
	return repositories.GetAllMaterial()
}

func GetMaterialByID(id int) (*models.Material, error) {
	return repositories.GetMaterialByID(id)
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

func UpdateMaterial(id int, data map[string]interface{}) (*models.Material, error){
	return repositories.UpdateMaterial(id, data)
}

func DeleteMaterial(id int) (*models.Material, error){
	return repositories.DeleteMaterial(id)
}
