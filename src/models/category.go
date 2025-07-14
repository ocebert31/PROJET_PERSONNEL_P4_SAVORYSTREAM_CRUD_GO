package models

import (
    "github.com/google/uuid"
	"time"
)

type Category struct { 
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name           string    `gorm:"size:50;not null;unique" json:"name"`
	CreatedAt      time.Time  `json:"created_at"`
    UpdatedAt      time.Time  `json:"updated_at"`

	Sauces  []Sauce `gorm:"foreignKey:CategoryID" json:"sauces"`
}
