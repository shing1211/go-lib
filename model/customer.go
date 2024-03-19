package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	CustID            string             `gorm:"column:CustID" bson:"custid" json:"custid,omitempty"`
	ClnID             string             `gorm:"column:ClnID" bson:"clnid" json:"clnid,omitempty"`
	SubACID           string             `gorm:"column:SubACID" bson:"subacid" json:"subacid,omitempty"`
	SrvSubACShortName string             `gorm:"column:SrvSubACShortName" bson:"srvsubacshortname" json:"srvsubacshortname,omitempty"`
	Holdings          []Holding          `gorm:"column:Holdings" bson:"holdings" json:"holdings"`
	CreatedAt         time.Time          `gorm:"column:CreatedAt" bson:"createdat" json:"createdat"`
	UpdatedAt         time.Time          `gorm:"column:UpdatedAt" bson:"updatedat" json:"updatedat"`
}
