package views

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums/schemecriteria"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type SchemeListView struct {
	Schemes []SchemeView `json:"schemes"`
}

type SchemeView struct {
	ID       uuid.UUID            `json:"id"`
	Name     string               `json:"name"`
	Criteria []schemeCriteriaView `json:"criteria"`
	Benefits []schemeBenefitView  `json:"benefits"`
}

type schemeCriteriaView struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type schemeBenefitView struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount float32   `json:"amount"`
}

func SchemeListViewFrom(schemes []models.Scheme) SchemeListView {
	schemeViews := make([]SchemeView, len(schemes))
	for i, s := range schemes {
		schemeViews[i] = SchemeView{
			ID:       s.ID,
			Name:     s.Name,
			Criteria: criteriaViewsFrom(s.Criteria),
			Benefits: benefitViewsFrom(s.Benefits),
		}
	}
	return SchemeListView{Schemes: schemeViews}
}

func criteriaViewsFrom(criteria []models.SchemeCriteria) []schemeCriteriaView {
	criteriaViews := make([]schemeCriteriaView, len(criteria))
	for i, c := range criteria {
		criteriaViews[i] = schemeCriteriaView{
			Key:   schemecriteria.KeyToString[c.CriteriaKey],
			Value: c.CriteriaValue,
		}
	}
	return criteriaViews
}

func benefitViewsFrom(benefits []models.SchemeBenefit) []schemeBenefitView {
	benefitViews := make([]schemeBenefitView, len(benefits))
	for i, b := range benefits {
		benefitViews[i] = schemeBenefitView{
			ID:     b.ID,
			Name:   b.Description,
			Amount: b.Amount,
		}
	}
	return benefitViews
}
