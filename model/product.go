package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	CreatedTimestamp time.Time      `gorm:"column:createdtimestamp;type:timestamp not null default CURRENT_TIMESTAMP" bson:"createdtimestamp" json:"createdtimestamp,omitempty"`
	UpdatedTimestamp time.Time      `gorm:"not null;column:updatedtimestamp;type:timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" bson:"updatedtimestamp" json:"updatedtimestamp,omitempty"`
	DeletedTimestamp gorm.DeletedAt `gorm:"null;index;column:deletedtimestamp;type:timestamp null" bson:"deletedtimestamp" json:"-"`
	CreatedBy        string         `gorm:"column:createdby;type:varchar(20) not null" bson:"createdby" json:"createdby,omitempty"`
	UpdatedBy        string         `gorm:"column:updatedby;type:varchar(20) not null" bson:"updatedby" json:"updatedby,omitempty"`
	DeletedBy        string         `gorm:"column:deletedby;type:varchar(20) null" bson:"deletedby" json:"deletedby,omitempty"`
	InstruID         string         `gorm:"primaryKey;column:instruid;type:varchar(7) not null" bson:"instruid" json:"instruid,omitempty"`
	PLCode           string         `gorm:"column:plcode;type:varchar(3) not null" bson:"plcode" json:"plcode,omitempty"`
	StockCode        string         `gorm:"column:stockcode;type:varchar(10) not null" bson:"stockcode" json:"stockcode,omitempty"`
	InstruShortName  string         `gorm:"column:instrushortname;type:varchar(40) not null" bson:"instrushortname" json:"instrushortname,omitempty"`
}

func (Product) TableName() string {
	return "product"
}
