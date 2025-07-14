package models

import (
    "github.com/google/uuid"
	"time"
)

type Stock struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Quantity  int       `gorm:"not null" json:"quantity"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
	SauceID   uuid.UUID `gorm:"type:uuid;not null" json:"sauce_id"`
}