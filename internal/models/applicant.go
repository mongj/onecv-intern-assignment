package models

import (
	"github.com/google/uuid"
)

type Applicant struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	PersonID uuid.UUID `gorm:"type:uuid;not null"`
	Person   *Person   `gorm:"->;<-:create"`
}
