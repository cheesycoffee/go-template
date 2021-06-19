package dao

import "time"

// RecordProperty of model
type RecordProperty struct {
	CreatedBy uint64    `json:"created_by" bson:"createdBy" gorm:"column:created_by;type:BIGINT(20)"`
	CreatedAt time.Time `json:"created_at" bson:"createdAt" gorm:"column:created_at;type:TIMESTAMP;default:NOW"`
	UpdatedBy uint64    `json:"updated_by" bson:"updatedBy" gorm:"column:updated_by;type:BIGINT(20)"`
	UpdatedAt time.Time `json:"updated_at" bson:"updatedAt" gorm:"column:updated_at;type:TIMESTAMP;default:NOW"`
}
