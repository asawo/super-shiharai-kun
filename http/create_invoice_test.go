package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/morikuni/failure"

	"github.com/asawo/api/errors"
	"github.com/asawo/api/service"
)

func Test_CreateInvoice_validation(t *testing.T) {
	t.Parallel()

	// Create test data
	now := time.Now()
	oneYearLater := now.AddDate(1, 0, 0)
	oneYearBefore := now.AddDate(-1, 0, 0)

	validBodyJson, err := json.Marshal(map[string]interface{}{
		"payment_amount":      1000,
		"service_provider_id": 1,
		"due_date":            oneYearLater,
	})
	if err != nil {
		t.Errorf("error marshalling json %v", err)
	}
	validBody := bytes.NewBuffer(validBodyJson)

	invalidDateJson, err := json.Marshal(map[string]interface{}{
		"payment_amount":      1000,
		"service_provider_id": 1,
		"due_date":            oneYearBefore,
	})
	if err != nil {
		t.Errorf("error marshalling json %v", err)
	}
	invalidDate := bytes.NewBuffer(invalidDateJson)

	invalidPaymentAmountJson, err := json.Marshal(map[string]interface{}{
		"payment_amount":      0,
		"service_provider_id": 1,
		"due_date":            oneYearLater,
	})
	if err != nil {
		t.Errorf("error marshalling json %v", err)
	}
	invalidPaymentAmount := bytes.NewBuffer(invalidPaymentAmountJson)

	testCases := []struct {
		name      string
		req       *http.Request
		want      *service.CreateInvoiceRequest
		error     bool
		errorCode failure.Code
	}{
		{
			name: "request_method_should_be_POST",
			req: &http.Request{
				Method: "GET",
				Header: map[string][]string{},
			},
			want:      nil,
			error:     true,
			errorCode: errors.MethodNotAllowed,
		},
		{
			name: "request_header_should_contain_basic_auth",
			req: &http.Request{
				Method: "POST",
				Header: map[string][]string{},
			},
			want:      nil,
			error:     true,
			errorCode: errors.Unauthorized,
		},
		{
			name: "request_header_should_have_empty_or_correct_content_type",
			req: &http.Request{
				Method: "POST",
				Header: map[string][]string{
					"Content-Type": {"multipart/form-data"},
				},
			},
			want:      nil,
			error:     true,
			errorCode: errors.InvalidRequest,
		},
		{
			name: "basic_auth_header_should_be_formatted_correctly",
			req: &http.Request{
				Method: "POST",
				Header: map[string][]string{
					"Content-Type":  {"application/json"},
					"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("wrong_format")))},
				},
			},
			want:      nil,
			error:     true,
			errorCode: errors.Unauthorized,
		},
		{
			name: "due_date_should_not_be_backdated",
			req: &http.Request{
				Method: "POST",
				Header: map[string][]string{
					"Content-Type":  {"application/json"},
					"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("username:password")))},
				},
				Body: io.NopCloser(invalidDate),
			},
			want:      nil,
			error:     true,
			errorCode: errors.InvalidRequest,
		},
		{
			name: "payment_amount_should_be_above_0",
			req: &http.Request{
				Method: "POST",
				Header: map[string][]string{
					"Content-Type":  {"application/json"},
					"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("username:password")))},
				},
				Body: io.NopCloser(invalidPaymentAmount),
			},
			want:      nil,
			error:     true,
			errorCode: errors.InvalidRequest,
		},
		{
			name: "validate_request_successfully",
			req: &http.Request{
				Method: "POST",
				Header: map[string][]string{
					"Content-Type":  {"application/json"},
					"Authorization": {fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("username:password")))},
				},
				Body: io.NopCloser(validBody),
			},
			want: &service.CreateInvoiceRequest{
				Email:             "username",
				Password:          "password",
				ServiceProviderId: 1,
				PaymentAmount:     1000,
				DueDate:           oneYearLater,
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

			got, err := validateCreateInvoiceRequest(tc.req)
			if !tc.error {
				if err != nil {
					t.Errorf("want nil but error: %v", err.GetCode())
				}

				if diff := cmp.Diff(got, tc.want); diff != "" {
					t.Errorf("diff: (-got +want)\n%s", diff)
				}
			} else {
				if err == nil {
					t.Error("want error but nil")
				}
				if !err.ErrorCodeIs(tc.errorCode) {
					t.Errorf("wrong type of error: got %v, want %v", err.GetCode(), tc.errorCode)
				}
			}
		})
	}
}
