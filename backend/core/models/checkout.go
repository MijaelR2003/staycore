package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckOut struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	BookingID     uuid.UUID `json:"booking_id" gorm:"type:uuid;not null"`
	CheckedOutAt  time.Time `json:"checked_out_at"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
}

func (c *CheckOut) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	c.CheckedOutAt = time.Now()
	return nil
}