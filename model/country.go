package model

import (
	"time"

	"gorm.io/gorm"
)

// Country model - `Country` table
type Country struct {
	Id          uint64         `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy   string         `gorm:"not null;type:varchar(20)"`
	UpdatedBy   string         `gorm:"not null;type:varchar(20)"`
	DeletedBy   string         `gorm:"null;type:varchar(20)" json:"DeletedBy,omitempty"`
	CountryCode string         `gorm:"unique;not null;type:varchar(2)"`
	CountryName string         `gorm:"not null;type:varchar(40)"`
	Cities      []City         `gorm:"foreignkey:IdCountry;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:",omitempty"`
}
