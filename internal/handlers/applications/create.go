package applications

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

const createHandlerName = "applications::create"

func HandleCreate(w http.ResponseWriter, r *http.Request) ([]byte, int, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, createHandlerName))
	}

	var params params.Application
	err = json.DecodeParams(r.Body, &params)
	if err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(err, fmt.Sprintf(handlers.ErrDecodeParams, createHandlerName))
	}

	applications := params.ToModel()

	if err = models.CreateApplications(db, applications); err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "failed to create application")
	}

	return []byte(`{"message": "Application created successfully"}`), http.StatusOK, nil
}
