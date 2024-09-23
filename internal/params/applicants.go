package params

import (
	"strings"
	"time"

	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	timeutils "github.com/mongj/gds-onecv-swe-assignment/internal/utils"
)

type DateOnly time.Time

// UnmarshalJSON parses a JSON string into a DateOnly type
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	// Trim trailing and leading quotes from the JSON string
	s := strings.Trim(string(b), `"`)

	parsedDate, err := time.Parse(timeutils.DateFormat, s)
	if err != nil {
		return err
	}
	*d = DateOnly(parsedDate)
	return nil
}

type PersonParams struct {
	Name               string                 `json:"name"`
	EmploymentStatus   enums.EmploymentStatus `json:"employment_status"`
	Sex                enums.Sex              `json:"sex"`
	DateOfBirth        DateOnly               `json:"date_of_birth"`
	MaritalStatus      enums.MaritalStatus    `json:"marital_status"`
	CurrentSchoolLevel *enums.SchoolLevel     `json:"current_school_level,omitempty"`
}

type RelativeParams struct {
	PersonParams
	Relation enums.Relation `json:"relation"`
}

type ApplicantParams struct {
	PersonParams
	Household []RelativeParams `json:"household"`
}

// ToModel converts ApplicantParams to a Person representing the applicant
// and a slice of Relatives representing the applicant's family members.
func (p *ApplicantParams) ToModel() (models.Person, []models.Relative) {
	applicant := models.Person{
		Name:               p.Name,
		DateOfBirth:        time.Time(p.DateOfBirth),
		Sex:                p.Sex,
		EmploymentStatus:   p.EmploymentStatus,
		MaritalStatus:      p.MaritalStatus,
		CurrentSchoolLevel: p.CurrentSchoolLevel,
	}
	relatives := make([]models.Relative, len(p.Household))
	for i, relative := range p.Household {
		relatives[i] = models.Relative{
			Person: models.Person{
				Name:               relative.Name,
				DateOfBirth:        time.Time(relative.DateOfBirth),
				Sex:                relative.Sex,
				EmploymentStatus:   relative.EmploymentStatus,
				MaritalStatus:      relative.MaritalStatus,
				CurrentSchoolLevel: relative.CurrentSchoolLevel,
			},
			Relation: relative.Relation,
		}
	}
	return applicant, relatives
}
