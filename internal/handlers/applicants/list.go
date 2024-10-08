package applicants

import (
	"fmt"
	"net/http"

	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/internal/json"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"github.com/mongj/gds-onecv-swe-assignment/internal/views"
	"github.com/pkg/errors"
)

const listHandlerName = "applicants::list"

func HandleList(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, listHandlerName))
	}

	applicants, err := models.ListApplicants(db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list applicants")
	}
	applicantListViews := make([]views.ApplicantViews, len(applicants))
	for i, a := range applicants {
		households, err := models.HouseholdMembersByPersonID(db, a.PersonID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read household")
		}
		applicantListViews[i] = views.ApplicantFrom(a, households)
	}

	data, err := json.EncodeView(applicantListViews)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrEncodeView, listHandlerName))
	}

	return data, nil
}
