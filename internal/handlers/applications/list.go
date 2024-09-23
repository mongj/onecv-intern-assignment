package applications

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

const listHandlerName = "applications::list"

func HandleList(w http.ResponseWriter, r *http.Request) ([]byte, int, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, listHandlerName))
	}

	applications, err := models.ListApplications(db)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to list applicants")
	}
	applicantListViews := views.ApplicationListViewFrom(applications)

	data, err := json.EncodeView(applicantListViews)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrEncodeView, listHandlerName))
	}

	return data, http.StatusOK, nil
}
