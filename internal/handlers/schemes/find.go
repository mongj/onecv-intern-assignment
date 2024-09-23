package schemes

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/api/exterror"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/internal/json"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"github.com/mongj/gds-onecv-swe-assignment/internal/views"
	"github.com/pkg/errors"
)

const findHandlerName = "schemes::findEligible"

func HandleFind(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, findHandlerName))
	}

	id, err := uuid.Parse(r.URL.Query().Get("applicant"))
	if err != nil {
		return nil, &exterror.BadRequest{Message: fmt.Sprintf("failed to parse applicant ID: %v", err)}
	}

	schemes, err := models.ListEligibleSchemes(db, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch eligible schemes from database")
	}
	schemeListView := views.SchemeListFrom(schemes)

	data, err := json.EncodeView(schemeListView)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrEncodeView, listHandlerName))
	}

	return data, nil
}
