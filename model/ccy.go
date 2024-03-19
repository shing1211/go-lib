package model

import (
	"time"

	"gorm.io/gorm"
)

// Ccy model - `Ccy` table
type Ccy struct {
	Id        uint64         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy string         `gorm:"not null;type:varchar(20)"`
	UpdatedBy string         `gorm:"not null;type:varchar(20)"`
	DeletedBy string         `gorm:"null;type:varchar(20)" json:"DeletedBy,omitempty"`
	CcyCode   string         `gorm:"unique;not null;type:varchar(3)"`
	CcyName   string         `gorm:"not null;type:varchar(40)"`
	DecPlc    int8           `gorm:"not null"`
}
