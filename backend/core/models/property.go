package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Property struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string     `json:"name" gorm:"not null"`
	Address   string     `json:"address"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Timezone  string     `json:"timezone" gorm:"default:'America/La_Paz'"`
	Active    bool       `json:"active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (p *Property) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}