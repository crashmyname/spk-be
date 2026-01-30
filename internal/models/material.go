package models

type Material struct {
	ID         uint   `json:"material_id" gorm:"primaryKey; column:material_id"`
	MoldNumber string `json:"mold_number"`
	LampName   string `json:"lamp_name"`
	ModelName  string `json:"model_name"`
	Type       string `json:"type"`
}

func (Material) TableName() string {
	return "material"
}
