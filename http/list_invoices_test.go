package http

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"

	"github.com/asawo/api/errors"
	"github.com/asawo/api/service"
)

func Test_ListInvoices_validation(t *testing.T) {
	t.Parallel()

	// Create test dates
	now := time.Now()
	oneYearLater := now.AddDate(1, 0, 0).Local().Format("2006-01-02")
	oneYearLaterParsed, err := time.Parse("2006-01-02", oneYearLater)
	if err != nil {
		t.Errorf("failed to parse date %v", err)
	}

	oneYearBefore := now.AddDate(-1, 0, 0).Local().Format("2006-01-02")
	oneYearBeforeParsed, err := time.Parse("2006-01-02", oneYearBefore)
	if err != nil {
		t.Errorf("failed to parse date %v", err)
	}

	testCases := []struct {
		name      string
		method    string
		header    map[string][]string
		url       string
		want      *service.ListInvoicesRequest
		error     bool
		errorCode failure.Code
	}{
		{
			name:      "request_method_should_be_GET",
			method:    "POST",
			header:    map[string][]string{},
			url:       "",
			want:      nil,
			error:     true,
			errorCode: errors.MethodNotAllowed,
		},
		{
			name:      "request_header_should_contain_basic_auth",
			method:    "GET",
			header:    map[string][]string{},
			url:       "",
			want:      nil,
			error:     true,
			errorCode: errors.Unauthorized,
		},
		{
			name:   "basic_auth_header_should_be_formatted_correctly",
			method: "GET",
			header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("wrong_format")))},
			},
			url:       "",
			want:      nil,
			error:     true,
			errorCode: errors.Unauthorized,
		},
		{
			name:   "start_date_should_not_be_before_end_date",
			method: "GET",
			header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("username:password")))},
			},
			url:       fmt.Sprintf("http://localhost:8080/v1/invoices?start_date=%s&end_date=%s", oneYearLater, oneYearBefore),
			want:      nil,
			error:     true,
			errorCode: errors.InvalidRequest,
		},
		{
			name:   "validate_request_successfully",
			method: "GET",
			header: map[string][]string{
				"Content-Type":  {"application/json"},
				"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("username:password")))},
			},
			url: fmt.Sprintf("http://localhost:8080/v1/invoices?start_date=%s&end_date=%s", oneYearBefore, oneYearLater),
			want: &service.ListInvoicesRequest{
				Email:     "username",
				Password:  "password",
				StartDate: oneYearBeforeParsed,
				EndDate:   oneYearLaterParsed,
			},
			error:     false,
			errorCode: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			req, err := http.NewRequest(tc.method, tc.url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
				return
			}

			req.Header = tc.header

			got, serr := validateListInvoicesRequest(req)
			if !tc.error {
				if err != nil {
					t.Errorf("want nil but error: %v", serr.GetCode())
				}

				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("diff: (-got +want)\n%s", diff)
				}
			} else {
				if serr == nil {
					t.Error("want error but nil")
				}
				if !serr.ErrorCodeIs(tc.errorCode) {
					t.Errorf("wrong type of error: got %v, want %v", serr.GetCode(), tc.errorCode)
				}
			}
		})
	}
}
