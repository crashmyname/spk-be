package models

import "time"

type Handover struct {
	ID                   uint       `json:"handover_id" gorm:"primaryKey;column:handover_id"`
	TID                  uint       `json:"ticket_id"`
	HandOver             string     `json:"handover"`
	Result               string     `json:"result"`
	DateHandover         *time.Time `json:"date_handover"`
	ApprovedMoldSect     string     `json:"approved_mold_sect"`
	DateApprovedMoldSect *time.Time `json:"date_approved_mold_sect"`
	ApprovedSect         string     `json:"approved_sect_by"`
	DateApprovedSect     string     `json:"date_approved_sect"`
	ApproveQc            string     `json:"approved_qc_by"`
	DateApprovedQC       *time.Time `json:"date_approved_qc"`
	CreatedAt            time.Time  `gorm:"column:created_at"`
	UpdatedAt            time.Time  `gorm:"column:updated_at"`
}
