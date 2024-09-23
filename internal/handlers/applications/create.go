package applications

import (
	"fmt"
	"net/http"

	"github.com/mongj/gds-onecv-swe-assignment/internal/api/exterror"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers"
	"github.com/mongj/gds-onecv-swe-assignment/internal/json"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"github.com/mongj/gds-onecv-swe-assignment/internal/params"
	"github.com/pkg/errors"
)

const createHandlerName = "applications::create"

func HandleCreate(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, createHandlerName))
	}

	var params params.Application
	err = json.DecodeParams(r.Body, &params)
	if err != nil {
		return nil, &exterror.BadRequest{Message: fmt.Sprintf("failed to decode request body: %v", err)}
	}

	applications := params.ToModel()

	if err = models.CreateApplications(db, applications); err != nil {
		return nil, errors.Wrap(err, "failed to create application")
	}

	return []byte(`{"message": "Application created successfully"}`), nil
}
