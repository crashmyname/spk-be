package models

type Ticket struct {
	ID         uint   `json:"ticket_id" gorm:"primaryKey;column:ticket_id"`
	NoOrder    string `json:"no_order"`
	DateCreate string `json:"date_create"`
	UID        uint   `json:"user_id" gorm:"secondaryKey"`
	Action     string `json:"action"`
	MaterialID string `json:"material_id" gorm:"secondaryKey"`
	LotShot    string `json:"lot_shot"`
	TotalShot  string `json:"total_shot"`
	SketchItem string `json:"skecth_item"`
	Options    string `json:"options"`
}
