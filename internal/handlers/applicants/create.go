package applicants

import (
	"fmt"
	"net/http"

	"github.com/mongj/gds-onecv-swe-assignment/internal/dataaccess/applicants"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/internal/json"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/params"
	"github.com/pkg/errors"
)

const createHandlerName = "applicants::create"

func HandleCreate(w http.ResponseWriter, r *http.Request) ([]byte, int, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, createHandlerName))
	}

	var params params.ApplicantParams
	err = json.DecodeParams(r.Body, &params)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrDecodeParams, createHandlerName))
	}
	applicant, relatives := params.ToModel()
	if err = applicants.Create(db, applicant, relatives); err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to create applicant")
	}

	return []byte(`{"message": "Applicant created successfully"}`), http.StatusOK, nil
}
