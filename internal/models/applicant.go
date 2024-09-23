package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Applicant struct {
	PersonID uuid.UUID `gorm:"primaryKey;type:uuid;not null"`
	Person   *Person   `gorm:"->;<-:create"`
}

func (a *Applicant) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func ListApplicants(db *gorm.DB) ([]Applicant, error) {
	var a []Applicant
	if err := db.Preload("Person").Find(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// ReadApplicant returns an applicant from the database given by the ID
func ReadApplicant(db *gorm.DB, id uuid.UUID) (*Applicant, error) {
	var a Applicant
	if err := db.Preload("Person").First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}
