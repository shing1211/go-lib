package model

import (
	"time"

	"gorm.io/gorm"
)

type InstruDepXsact struct {
	CreatedTimestamp time.Time      `gorm:"column:createdtimestamp;type:timestamp not null default CURRENT_TIMESTAMP" bson:"createdtimestamp" json:"createdtimestamp,omitempty"`
	UpdatedTimestamp time.Time      `gorm:"not null;column:updatedtimestamp;type:timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" bson:"updatedtimestamp" json:"updatedtimestamp,omitempty"`
	DeletedTimestamp gorm.DeletedAt `gorm:"null;index;column:deletedtimestamp;type:timestamp null" bson:"deletedtimestamp" json:"-"`
	CreatedBy        string         `gorm:"column:createdby;type:varchar(20) not null" bson:"createdby" json:"createdby,omitempty"`
	UpdatedBy        string         `gorm:"column:updatedby;type:varchar(20) not null" bson:"updatedby" json:"updatedby,omitempty"`
	DeletedBy        string         `gorm:"column:deletedby;type:varchar(20) null" bson:"deletedby" json:"deletedby,omitempty"`
	ClnID            string         `gorm:"primaryKey;column:clnid;type:varchar(7) not null" bson:"clnid" json:"clnid,omitempty"`
	SubACID          string         `gorm:"primaryKey;column:subacid;type:varchar(4) not null" bson:"subacid" json:"subacid,omitempty"`
	InstruID         string         `gorm:"primaryKey;column:instruID;type:varchar(7) not null" bson:"instruid" json:"instruid,omitempty"`
	TotalQty         float64        `gorm:"not null;column:totalqty" bson:"totalqty" json:"totalqty,omitempty"`
	Product          Product        `gorm:"foreignKey:InstruID;references:InstruID" bson:"product" json:"product,omitempty"`
}

func (InstruDepXsact) TableName() string {
	return "instrudepxsact"
}
