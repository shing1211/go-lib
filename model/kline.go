package model

import (
	"time"

	"gorm.io/gorm"
)

// WsAggTradeEvent define websocket aggregate trade event
type Kline struct {
	Id                   uint64         `gorm:"primaryKey"`
	CreatedAt            time.Time      `gorm:"not null"`
	UpdatedAt            time.Time      `gorm:"not null"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy            string         `gorm:"not null;type:varchar(20)"`
	UpdatedBy            string         `gorm:"not null;type:varchar(20)"`
	DeletedBy            string         `gorm:"null;type:varchar(20)" json:"DeletedBy,omitempty"`
	Time                 int64
	Symbol               string `gorm:"index"`
	StartTime            int64  `gorm:"index"`
	EndTime              int64
	Interval             string `gorm:"index"`
	FirstTradeID         int64
	LastTradeID          int64
	Open                 string
	Close                string
	High                 string
	Low                  string
	Volume               string
	TradeNum             int64
	IsFinal              bool `gorm:"index"`
	QuoteVolume          string
	ActiveBuyVolume      string
	ActiveBuyQuoteVolume string
}
