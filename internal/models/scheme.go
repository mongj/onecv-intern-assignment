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

	hadChldQuery := db.
		Table("households h").
		Select("1").
		Where(fmt.Sprintf("h.person_id = a.person_id AND h.relation = '%s'", enums.RelationChild))

	chldSchLvlQuery := db.
		Table("people r").
		Select("ARRAY_AGG(DISTINCT r.current_school_level::TEXT)").
		Joins("JOIN households h ON r.id = h.relative_id").
		Where("h.person_id = a.person_id")

	applicantInfoQuery := db.
		Table("applicants a").
		Select(
			"p.*, EXISTS(?) AS has_children, (?) AS children_school_levels",
			hadChldQuery, chldSchLvlQuery,
		).
		Joins("JOIN people p ON a.person_id = p.id").
		Where("p.id = ?", id)

	allCriteriaQuery := db.
		Table("scheme_criteria sc").
		Select("sc.scheme_id, COUNT(sc.scheme_id) AS total_cnt").
		Group("sc.scheme_id")

	fulfilledCriteriaQuery := db.
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
		), applicantInfoQuery,
		).
		Group("sc.scheme_id")

	eligibleSchemeIDsQuery := db.
		Table("(?) AS fc", fulfilledCriteriaQuery).
		Select("fc.scheme_id").
		Joins("JOIN (?) AS ac ON fc.scheme_id = ac.scheme_id", allCriteriaQuery).
		Where("fc.fulfilled_cnt = ac.total_cnt")

	err := db.
		Model(&Scheme{}).
		Preload("Benefits").
		Preload("Criteria").
		Joins("JOIN (?) AS es ON schemes.id = es.scheme_id", eligibleSchemeIDsQuery).
		Find(&schemes).
		Error
	if err != nil {
		return nil, err
	}

	return schemes, nil
}
