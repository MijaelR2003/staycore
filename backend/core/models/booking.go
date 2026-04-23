package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	RoomID        uuid.UUID `json:"room_id" gorm:"type:uuid;not null"`
	GuestID       uuid.UUID `json:"guest_id" gorm:"type:uuid;not null"`
	Status        string    `json:"status" gorm:"default:'pending'"`
	CheckInDate   time.Time `json:"check_in_date" gorm:"not null"`
	CheckOutDate  time.Time `json:"check_out_date" gorm:"not null"`
	GuestCount    int       `json:"guest_count" gorm:"default:1"`
	Source        string    `json:"source" gorm:"default:'direct'"`
	ExternalID    string    `json:"external_id"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}