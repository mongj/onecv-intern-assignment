package params

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
)

type Application struct {
	ID        uuid.UUID   `json:"id"`
	SchemeIDs []uuid.UUID `json:"schemeIds"`
}

func (a *Application) ToModel() []models.Application {
	applications := make([]models.Application, len(a.SchemeIDs))
	for i, sid := range a.SchemeIDs {
		applications[i] = models.Application{
			ApplicantID:       a.ID,
			SchemeID:          sid,
			ApplicationStatus: enums.ApplicationStatusPending, // Default to pending
		}
	}
	return applications
}
