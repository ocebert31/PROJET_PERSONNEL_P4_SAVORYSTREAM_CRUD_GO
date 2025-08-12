package models

import (
    "github.com/google/uuid"
	"time"
)

type Stock struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Quantity  int       `gorm:"not null" json:"quantity"` 

    SauceID  uuid.UUID `gorm:"type:uuid;not null;unique" json:"-"`
    Sauce   Sauce     `gorm:"foreignKey:SauceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sauce"`

    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}