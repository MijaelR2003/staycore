package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	PropertyID uuid.UUID `json:"property_id" gorm:"type:uuid;not null"`
	Number     string    `json:"number" gorm:"not null"`
	Type       string    `json:"type" gorm:"not null"`
	Capacity   int       `json:"capacity" gorm:"default:1"`
	Floor      int       `json:"floor"`
	Price      float64   `json:"price"`
	Status     string    `json:"status" gorm:"default:'available'"`
	Active     bool      `json:"active" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (r *Room) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}