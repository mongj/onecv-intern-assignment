package views

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type ApplicationListView struct {
	Applications []Application `json:"applications"`
}

type Application struct {
	ID                uuid.UUID               `json:"id"`
	ApplicantID       uuid.UUID               `json:"applicant_id"`
	SchemeID          uuid.UUID               `json:"scheme_id"`
	ApplicationStatus enums.ApplicationStatus `json:"application_status"`
}

func ApplicationListViewFrom(applications []models.Application) ApplicationListView {
	applicationViews := make([]Application, len(applications))
	for i, a := range applications {
		applicationViews[i] = Application{
			ID:                a.ID,
			ApplicantID:       a.ApplicantID,
			SchemeID:          a.SchemeID,
			ApplicationStatus: a.ApplicationStatus,
		}
	}
	return ApplicationListView{Applications: applicationViews}
}
