package models

type DetailTicket struct {
	ID             uint   `json:"detail_id" gorm:"primaryKey; column:detail_id"`
	TID            uint   `json:"ticket_id"`
	DetailItem     string `json:"detail_item"`
	RepairReq      string `json:"repair_req"`
	DateRepair     string `json:"date_repair"`
	RepairBy       string `json:"repair_by"`
	TotalHoursPlan int    `json:"total_hours_plan"`
	ActRepair      string `json:"act_repair"`
	DateAct        string `json:"date_act"`
	ActBy          string `json:"act_by"`
	TotalHoursAct  string `json:"total_hours_act"`
}
