package models

import (
	"github.com/google/uuid"
)

type Applicant struct {
	PersonID uuid.UUID `gorm:"type:uuid;not null"`
	Person   *Person   `gorm:"->;<-:create"`
}
