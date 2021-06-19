package dto

import "time"

// RecordProperty of model
type RecordProperty struct {
	CreatedBy uint64    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy uint64    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}
