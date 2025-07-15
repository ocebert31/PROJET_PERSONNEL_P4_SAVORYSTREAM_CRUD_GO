package models

import (
    "github.com/google/uuid"
    "time"
)

type Sauce struct {
    ID             uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
    Name           string        `gorm:"size:50;not null;unique" json:"name"`
    Description    string        `gorm:"type:text;not null" json:"description"`
    Characteristic *string       `gorm:"size:255" json:"characteristic,omitempty"`

    IsAvailable    bool          `gorm:"default:true" json:"is_available"`

    CategoryID     uuid.UUID     `gorm:"type:uuid;not null;index" json:"category_id"`
    Category       *Category     `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"category"`

    Stock          *Stock        `gorm:"foreignKey:SauceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"stock,omitempty"`

    Ingredients    []Ingredient  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"ingredients"`
    Conditionings  []Conditioning `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"conditionings"`

    CreatedAt      time.Time     `json:"created_at"`
    UpdatedAt      time.Time     `json:"updated_at"`
}
