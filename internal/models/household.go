package models

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"gorm.io/gorm"
)

type Household struct {
	ID         int            `gorm:"primaryKey"`
	PersonID   uuid.UUID      `gorm:"type:uuid;not null"`
	Person     Person         `gorm:"foreignKey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RelativeID uuid.UUID      `gorm:"type:uuid;not null"`
	Relative   Person         `gorm:"foreignKey:RelativeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Relation   enums.Relation `gorm:"type:relation;not null"`
}

func (h *Household) Create(db *gorm.DB) error {
	return db.Create(&h).Error
}

func HouseholdMembersByPersonID(db *gorm.DB, personID uuid.UUID) ([]Household, error) {
	var h []Household
	if err := db.Preload("Relative").Where("person_id = ?", personID).Find(&h).Error; err != nil {
		return []Household{}, err
	}
	return h, nil
}
