package models

import (
    "github.com/google/uuid"
	"time"
)

type Conditioning struct { 
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Volume         string    `gorm:"size:20;not null" json:"volume"`
	Price    	   float64   `gorm:"type:decimal(10,2);not null" json:"price"`

	SauceID uuid.UUID `gorm:"type:uuid;not null" json:"sauce_id"`
	Sauce   Sauce     `gorm:"foreignKey:SauceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sauce"`

	CreatedAt      time.Time  `json:"created_at"`
    UpdatedAt      time.Time  `json:"updated_at"`
}


