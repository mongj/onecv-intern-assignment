package views

import (
	"time"

	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type PersonView struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	EmploymentStatus   enums.EmploymentStatus `json:"employment_status"`
	Sex                enums.Sex              `json:"sex"`
	DateOfBirth        time.Time              `json:"date_of_birth"`
	MaritalStatus      enums.MaritalStatus    `json:"marital_status"`
	CurrentSchoolLevel *enums.SchoolLevel     `json:"current_school_level,omitempty"`
}

type RelativeView struct {
	PersonView
	Relation enums.Relation `json:"relation"`
}

type ApplicantViews struct {
	PersonView
	Household []RelativeView `json:"household"`
}

func ApplicantViewFrom(applicant models.Applicant, household []models.Household) ApplicantViews {
	return ApplicantViews{
		PersonView: PersonView{
			ID:                 applicant.Person.ID.String(),
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

func relativeViewsFrom(household []models.Household) []RelativeView {
	relativeViews := make([]RelativeView, len(household))
	for i, h := range household {
		relativeViews[i] = RelativeView{
			PersonView: PersonView{
				ID:                 h.Relative.ID.String(),
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
