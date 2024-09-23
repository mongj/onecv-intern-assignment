package models

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums/schemecriteria"
)

type Scheme struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name     string    `gorm:"not null"`
	Benefits []SchemeBenefit
	Criteria []SchemeCriteria
}

type SchemeBenefit struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SchemeID    uuid.UUID `gorm:"type:uuid"`
	Description string    `gorm:"not null"`
	Amount      float32   `gorm:"not null;type:decimal(12,2)"`
}

type SchemeCriteria struct {
	ID          int                `gorm:"primaryKey"`
	SchemeID    uuid.UUID          `gorm:"type:uuid"`
	CriteriaKey schemecriteria.Key `gorm:"not null"`
	// Type of the value is inferred from the criteria key when the value is used
	CriteriaValue string `gorm:"not null"`
}
