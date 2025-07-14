package models

import (
    "github.com/google/uuid"
    "time"
)

type Sauce struct {
    ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name           string     `gorm:"size:50;not null;unique" json:"name"`
    Description    string     `gorm:"type:text;not null" json:"description"`
    Characteristic string     `gorm:"size:255" json:"characteristic"`
    IsAvailable    bool       `gorm:"default:true" json:"is_available"`
    CategoryID     uuid.UUID  `gorm:"type:uuid" json:"category_id"`
	Stock          Stock `gorm:"foreignKey:SauceID" json:"stock"`
    CreatedAt      time.Time  `json:"created_at"`
    UpdatedAt      time.Time  `json:"updated_at"`

	Ingredients   []Ingredient `gorm:"foreignKey:SauceID" json:"ingredients"`
	Conditionings []Conditioning `gorm:"foreignKey:SauceID" json:"conditionings"`
}
