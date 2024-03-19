package model

import (
	"time"

	"gorm.io/gorm"
)

// WsAggTradeEvent define websocket aggregate trade event
type Aggtrade struct {
	Id                    uint64         `gorm:"primaryKey"`
	CreatedAt             time.Time      `gorm:"not null"`
	UpdatedAt             time.Time      `gorm:"not null"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy             string         `gorm:"not null;type:varchar(20)"`
	UpdatedBy             string         `gorm:"not null;type:varchar(20)"`
	DeletedBy             string         `gorm:"null;type:varchar(20)" json:"DeletedBy,omitempty"`
	Time                  int64
	Symbol                string `gorm:"index"`
	AggTradeID            int64
	Price                 string
	Quantity              string
	FirstBreakdownTradeID int64
	LastBreakdownTradeID  int64
	TradeTime             int64
	IsBuyerMaker          bool
	Placeholder           bool
}
