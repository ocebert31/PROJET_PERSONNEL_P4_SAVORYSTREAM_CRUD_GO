package models

import (
    "github.com/google/uuid"
	"time"
)

type Ingredient struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name      string    `gorm:"size:100;not null" json:"name"`
	Quantity  string    `gorm:"size:100;not null" json:"quantity"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
	SauceID   uuid.UUID  `gorm:"type:uuid" json:"sauce_id"`
}