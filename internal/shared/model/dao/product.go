package dao

import "github.com/cheesycoffee/go-template/internal/shared/constant"

// Product model
type Product struct {
	BID         string `json:"-" bson:"_id" gorm:"-"`
	ID          uint64 `json:"id" bson:"id" gorm:"column:id;primaryKey;autoIncrement;type:bigint(20)"`
	Description string `json:"description" bson:"description" gorm:"column:description;type:text"`
	RecordProperty
}

// TableName : gorm table name implementation
func (Product) TableName() string {
	return constant.TableProduct
}
