package applicants

import (
	"fmt"
	"net/http"

	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/internal/json"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
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
		return nil, http.StatusBadRequest, errors.Wrap(err, fmt.Sprintf(handlers.ErrDecodeParams, createHandlerName))
	}
	person, relatives := params.ToModel()
	// Create the applicant
	a := models.Applicant{
		PersonID: person.ID,
		Person:   &person,
	}
	if err = a.Create(db); err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to create applicant")
	}

	// Create and link the relatives
	for _, r := range relatives {
		h := models.Household{
			PersonID: person.ID,
			Relative: r.Person,
			Relation: r.Relation,
		}
		if err = h.Create(db); err != nil {
			return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to create household")
		}
	}

	return []byte(`{"message": "Applicant created successfully"}`), http.StatusOK, nil
}
