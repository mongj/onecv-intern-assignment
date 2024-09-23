package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
)

type Person struct {
	ID                 uuid.UUID              `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name               string                 `gorm:"not null"`
	DateOfBirth        time.Time              `gorm:"not null"`
	Sex                enums.Sex              `gorm:"type:sex;not null"`
	EmploymentStatus   enums.EmploymentStatus `gorm:"type:employment_status;not null"`
	MaritalStatus      enums.MaritalStatus    `gorm:"type:marital_status;not null"`
	CurrentSchoolLevel *enums.SchoolLevel     `gorm:"type:school_level"`
}

type Relative struct {
	Person
	Relation enums.Relation `gorm:"type:relation;not null"`
}

func (Person) TableName() string {
	return "people"
}
