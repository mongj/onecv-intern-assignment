package schemes

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

const listHandlerName = "schemes::list"

func HandleList(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	db, err := middleware.GetDB(r)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrGetDB, listHandlerName))
	}

	schemes, err := models.ListSchemes(db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch schemes from database")
	}
	schemeListView := views.SchemeListViewFrom(schemes)

	data, err := json.EncodeView(schemeListView)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(handlers.ErrEncodeView, listHandlerName))
	}

	return data, nil
}
