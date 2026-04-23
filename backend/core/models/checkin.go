package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckIn struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	BookingID    uuid.UUID `json:"booking_id" gorm:"type:uuid;not null"`
	CheckedInAt  time.Time `json:"checked_in_at"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

func (c *CheckIn) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	c.CheckedInAt = time.Now()
	return nil
}