package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guest struct {
	ID              uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName       string     `json:"first_name" gorm:"not null"`
	LastName        string     `json:"last_name" gorm:"not null"`
	DocumentType    string     `json:"document_type" gorm:"not null"`
	DocumentNumber  string     `json:"document_number" gorm:"not null"`
	DocumentCountry string     `json:"document_country"`
	Nationality     string     `json:"nationality"`
	BirthDate       *time.Time `json:"birth_date"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (g *Guest) BeforeCreate(tx *gorm.DB) error {
	g.ID = uuid.New()
	return nil
}