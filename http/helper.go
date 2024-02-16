package http

import (
	"net/http"
	"time"

	"github.com/morikuni/failure"

	"github.com/asawo/api/errors"
)

func GetQueryParamsForListInvoices(r *http.Request) (time.Time, time.Time, errors.ServiceError) {
	// Parse query parameters
	startDateParam := r.URL.Query().Get("start_date")
	endDateParam := r.URL.Query().Get("end_date")

	// Parse start_date parameter
	var startDate time.Time
	if startDateParam != "" {
		var err error
		startDate, err = time.Parse("2006-01-02", startDateParam)
		if err != nil {
			errors.NewFromError(err, errors.InvalidRequest, failure.Message("Invalid startDate format. Please use YYYY-MM-DD"))
		}
	}

	// Parse end_date parameter
	var endDate time.Time
	if endDateParam != "" {
		var err error
		endDate, err = time.Parse("2006-01-02", endDateParam)
		if err != nil {
			errors.NewFromError(err, errors.InvalidRequest, failure.Message("Invalid due_date format. Please use YYYY-MM-DD"))
		}
	}

	return startDate, endDate, nil
}
