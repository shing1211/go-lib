package model

import (
	"time"

	"gorm.io/gorm"
)

// City model - `City` table
type City struct {
	Id        uint64         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy string         `gorm:"not null;type:varchar(20)"`
	UpdatedBy string         `gorm:"not null;type:varchar(20)"`
	DeletedBy string         `gorm:"null;type:varchar(20)" json:"DeletedBy,omitempty"`
	CityCode  string         `gorm:"unique;not null;type:varchar(2)"`
	CityName  string         `gorm:"not null;type:varchar(40)"`
	IdCountry uint64         `gorm:"not null"`
}
