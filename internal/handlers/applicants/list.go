package applicants

import (
	"fmt"
	"net/http"

	"github.com/mongj/gds-onecv-swe-assignment/internal/dataaccess/applicants"
	"github.com/mongj/gds-onecv-swe-assignment/internal/dataaccess/households"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/internal/json"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/views"
	"github.com/pkg/errors"
)

const listHandlerName = "applicants::list"

func HandleList(w http.ResponseWriter, r *http.Request) ([]byte, int, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, listHandlerName))
	}

	applicants, err := applicants.List(db)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to list applicants")
	}
	applicantListViews := make([]views.ApplicantViews, len(applicants))
	for i, a := range applicants {
		// Get relative view
		households, err := households.ReadByPersonID(db, a.PersonID)
		if err != nil {
			return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to read household")
		}
		applicantListViews[i] = views.ApplicantViewFrom(a, households)
	}

	data, err := json.EncodeView(applicantListViews)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrEncodeView, listHandlerName))
	}

	return data, http.StatusOK, nil
}
