package views

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type ApplicationList struct {
	Applications []Application `json:"applications"`
}

type Application struct {
	ID                uuid.UUID               `json:"id"`
	ApplicantID       uuid.UUID               `json:"applicant_id"`
	SchemeID          uuid.UUID               `json:"scheme_id"`
	ApplicationStatus enums.ApplicationStatus `json:"application_status"`
}

func ApplicationListFrom(applications []models.Application) ApplicationList {
	applicationViews := make([]Application, len(applications))
	for i, a := range applications {
		applicationViews[i] = Application{
			ID:                a.ID,
			ApplicantID:       a.ApplicantID,
			SchemeID:          a.SchemeID,
			ApplicationStatus: a.ApplicationStatus,
		}
	}
	return ApplicationList{Applications: applicationViews}
}
