package views

import (
	"time"

	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type Person struct {
	ID                 uuid.UUID              `json:"id"`
	Name               string                 `json:"name"`
	EmploymentStatus   enums.EmploymentStatus `json:"employment_status"`
	Sex                enums.Sex              `json:"sex"`
	DateOfBirth        time.Time              `json:"date_of_birth"`
	MaritalStatus      enums.MaritalStatus    `json:"marital_status"`
	CurrentSchoolLevel *enums.SchoolLevel     `json:"current_school_level,omitempty"`
}

type Relative struct {
	Person
	Relation enums.Relation `json:"relation"`
}

type ApplicantViews struct {
	Person
	Household []Relative `json:"household"`
}

func ApplicantFrom(applicant models.Applicant, household []models.Household) ApplicantViews {
	return ApplicantViews{
		Person: Person{
			ID:                 applicant.Person.ID,
			Name:               applicant.Person.Name,
			EmploymentStatus:   applicant.Person.EmploymentStatus,
			Sex:                applicant.Person.Sex,
			DateOfBirth:        applicant.Person.DateOfBirth,
			MaritalStatus:      applicant.Person.MaritalStatus,
			CurrentSchoolLevel: applicant.Person.CurrentSchoolLevel,
		},
		Household: relativeViewsFrom(household),
	}
}

func relativeViewsFrom(household []models.Household) []Relative {
	relativeViews := make([]Relative, len(household))
	for i, h := range household {
		relativeViews[i] = Relative{
			Person: Person{
				ID:                 h.Relative.ID,
				Name:               h.Relative.Name,
				EmploymentStatus:   h.Relative.EmploymentStatus,
				Sex:                h.Relative.Sex,
				DateOfBirth:        h.Relative.DateOfBirth,
				MaritalStatus:      h.Relative.MaritalStatus,
				CurrentSchoolLevel: h.Relative.CurrentSchoolLevel,
			},
			Relation: h.Relation,
		}
	}
	return relativeViews
}
