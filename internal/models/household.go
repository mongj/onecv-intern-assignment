package models

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
)

type Household struct {
	ID         int            `gorm:"primaryKey"`
	PersonID   uuid.UUID      `gorm:"type:uuid;not null"`
	RelativeID uuid.UUID      `gorm:"type:uuid;not null"`
	Relation   enums.Relation `gorm:"type:relation;not null"`

	Person   Person `gorm:"foreignKey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Relative Person `gorm:"foreignKey:RelativeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
