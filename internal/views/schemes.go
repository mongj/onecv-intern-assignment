package views

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums/schemecriteria"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type SchemeList struct {
	Schemes []Scheme `json:"schemes"`
}

type Scheme struct {
	ID       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	Criteria []schemeCriteria `json:"criteria"`
	Benefits []schemeBenefit  `json:"benefits"`
}

type schemeCriteria struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type schemeBenefit struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount float32   `json:"amount"`
}

func SchemeListFrom(schemes []models.Scheme) SchemeList {
	schemeViews := make([]Scheme, len(schemes))
	for i, s := range schemes {
		schemeViews[i] = Scheme{
			ID:       s.ID,
			Name:     s.Name,
			Criteria: criteriaViewsFrom(s.Criteria),
			Benefits: benefitViewsFrom(s.Benefits),
		}
	}
	return SchemeList{Schemes: schemeViews}
}

func criteriaViewsFrom(criteria []models.SchemeCriteria) []schemeCriteria {
	criteriaViews := make([]schemeCriteria, len(criteria))
	for i, c := range criteria {
		criteriaViews[i] = schemeCriteria{
			Key:   schemecriteria.KeyToString[c.CriteriaKey],
			Value: c.CriteriaValue,
		}
	}
	return criteriaViews
}

func benefitViewsFrom(benefits []models.SchemeBenefit) []schemeBenefit {
	benefitViews := make([]schemeBenefit, len(benefits))
	for i, b := range benefits {
		benefitViews[i] = schemeBenefit{
			ID:     b.ID,
			Name:   b.Description,
			Amount: b.Amount,
		}
	}
	return benefitViews
}
