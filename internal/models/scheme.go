package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums/schemecriteria"
	"gorm.io/gorm"
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

// isEligible checks if a person is eligible for a scheme based on the scheme's criteria
func (s *Scheme) isEligible(db *gorm.DB, id uuid.UUID) (bool, error) {
	sc := s.Criteria
	var p Person
	if err := db.Model(&Person{}).First(&p, id).Error; err != nil {
		return false, err
	}

	for _, c := range sc {
		switch c.CriteriaKey {
		case schemecriteria.EmploymentStatus:
			if p.EmploymentStatus != enums.EmploymentStatus(c.CriteriaValue) {
				return false, nil
			}
		case schemecriteria.MaritalStatus:
			if p.MaritalStatus != enums.MaritalStatus(c.CriteriaValue) {
				return false, nil
			}
		case schemecriteria.HasChildren:
			var cnt int64
			err := db.
				Model(&Household{}).
				Where("person_id = ? AND relation = ?", id, enums.RelationChild).
				Count(&cnt).
				Error
			if err != nil {
				return false, err
			}
			if c.CriteriaValue == "true" && cnt == 0 || c.CriteriaValue == "false" && cnt > 0 {
				return false, nil
			}
		case schemecriteria.ChildrenSchoolLevel:
			var cnt int64
			err := db.
				Model(&Household{}).
				Joins("JOIN people ON households.relative_id = people.id").
				Where("households.person_id = ? AND households.relation = ?", id, enums.RelationChild).
				Where("people.current_school_level = ?", enums.SchoolLevel(c.CriteriaValue)).
				Count(&cnt).
				Error
			if err != nil {
				return false, err
			}
			if cnt == 0 {
				return false, nil
			}
		}
	}
	return true, nil
}

// ReadScheme returns a scheme from the database given by the ID
func ReadScheme(db *gorm.DB, id uuid.UUID) (*Scheme, error) {
	var s *Scheme
	if err := db.Preload("Benefits").Preload("Criteria").First(&s, id).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func ListSchemes(db *gorm.DB) ([]Scheme, error) {
	var s []Scheme
	if err := db.Preload("Benefits").Preload("Criteria").Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func ListEligibleSchemes(db *gorm.DB, id uuid.UUID) ([]Scheme, error) {
	var schemes []Scheme

	hasChldQry := db.
		Table("households h").
		Select("1").
		Where(fmt.Sprintf("h.person_id = a.person_id AND h.relation = '%s'", enums.RelationChild))

	chldSchLvlQry := db.
		Table("people r").
		Select("ARRAY_AGG(DISTINCT r.current_school_level::TEXT)").
		Joins("JOIN households h ON r.id = h.relative_id").
		Where("h.person_id = a.person_id")

	applicantInfoQry := db.
		Table("applicants a").
		Select(
			"p.*, EXISTS(?) AS has_children, (?) AS children_school_levels",
			hasChldQry, chldSchLvlQry,
		).
		Joins("JOIN people p ON a.person_id = p.id").
		Where("p.id = ?", id)

	allCtrQry := db.
		Table("scheme_criteria sc").
		Select("sc.scheme_id, COUNT(sc.scheme_id) AS total_cnt").
		Group("sc.scheme_id")

	fulfilledCtrQry := db.
		Table("scheme_criteria sc").
		Select("sc.scheme_id, COUNT(sc.scheme_id) AS fulfilled_cnt").
		Joins(fmt.Sprintf(`
		JOIN (?) AS ai ON (
			CASE
				WHEN sc.criteria_key = %d THEN sc.criteria_value = ai.employment_status::TEXT
				WHEN sc.criteria_key = %d THEN sc.criteria_value = ai.marital_status::TEXT
				WHEN sc.criteria_key = %d THEN sc.criteria_value = ai.has_children::TEXT
				WHEN sc.criteria_key = %d THEN sc.criteria_value = ANY(ai.children_school_levels)
				ELSE FALSE
			END)`,
			schemecriteria.EmploymentStatus,
			schemecriteria.MaritalStatus,
			schemecriteria.HasChildren,
			schemecriteria.ChildrenSchoolLevel,
		), applicantInfoQry,
		).
		Group("sc.scheme_id")

	eligibleSchemeIDsQry := db.
		Table("(?) AS fc", fulfilledCtrQry).
		Select("fc.scheme_id").
		Joins("JOIN (?) AS ac ON fc.scheme_id = ac.scheme_id", allCtrQry).
		Where("fc.fulfilled_cnt = ac.total_cnt")

	err := db.
		Model(&Scheme{}).
		Preload("Benefits").
		Preload("Criteria").
		Joins("JOIN (?) AS es ON schemes.id = es.scheme_id", eligibleSchemeIDsQry).
		Find(&schemes).
		Error
	if err != nil {
		return nil, err
	}

	return schemes, nil
}
