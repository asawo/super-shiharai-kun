package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/morikuni/failure"

	"github.com/asawo/api/auth"
	"github.com/asawo/api/db"
	"github.com/asawo/api/db/model"
	"github.com/asawo/api/errors"
)

func Test_ListInvoices(t *testing.T) {
	t.Parallel()

	now := time.Now()
	oneYearLater := now.AddDate(1, 0, 0)
	// oneYearBefore := now.AddDate(-1, 0, 0)

	testCases := []struct {
		name      string
		dbFactory func(mockCtrl *gomock.Controller) db.DB
		req       *ListInvoicesRequest
		want      []*model.Invoice
		error     bool
		errorCode failure.Code
	}{
		{
			name: "auth_should_be_valid",
			dbFactory: func(mockCtrl *gomock.Controller) db.DB {
				dbService := db.NewMockDB(mockCtrl)

				pass, err := auth.HashPassword("wrongpassword")
				if err != nil {
					t.Errorf("failed to hash password")
				}

				dbService.EXPECT().GetUserByEmail(gomock.Any(), "test@test.com").
					Return(&model.User{
						CompanyID: 1,
						Name:      "test",
						Email:     "test@test.com",
						Password:  pass,
					}, nil).AnyTimes().AnyTimes()

				return dbService
			},
			req: &ListInvoicesRequest{
				Email:     "test@test.com",
				Password:  "testpassword",
				StartDate: now,
				EndDate:   oneYearLater, // 1 year in the future
			},
			want:      nil,
			error:     true,
			errorCode: errors.Unauthorized,
		},
		{
			name: "successfully_list_invoices",
			dbFactory: func(mockCtrl *gomock.Controller) db.DB {
				dbService := db.NewMockDB(mockCtrl)

				pass, err := auth.HashPassword("testpassword")
				if err != nil {
					t.Errorf("failed to hash password")
				}

				dbService.EXPECT().GetUserByEmail(gomock.Any(), "test@test.com").
					Return(&model.User{
						CompanyID: 1,
						Name:      "test",
						Email:     "test@test.com",
						Password:  pass,
					}, nil)

				dbService.EXPECT().ListInvoices(gomock.Any(), uint(1), now, oneYearLater).
					Return([]*model.Invoice{
						{
							PaymentAmount:     100.00,
							Commission:        4.00,
							CommissionRate:    0.04,
							Tax:               0.4,
							TaxRate:           0.1,
							Amount:            104.40,
							DueDate:           oneYearLater,
							Status:            "OUTSTANDING",
							CompanyID:         1,
							ServiceProviderID: 1,
						},
						{
							PaymentAmount:     100.00,
							Commission:        4.00,
							CommissionRate:    0.04,
							Tax:               0.4,
							TaxRate:           0.1,
							Amount:            104.40,
							DueDate:           oneYearLater,
							Status:            "OUTSTANDING",
							CompanyID:         1,
							ServiceProviderID: 2,
						},
					}, nil)

				return dbService
			},
			req: &ListInvoicesRequest{
				Email:     "test@test.com",
				Password:  "testpassword",
				StartDate: now,
				EndDate:   oneYearLater, // 1 year in the future
			},
			want: []*model.Invoice{
				{
					PaymentAmount:     100.00,
					Commission:        4.00,
					CommissionRate:    0.04,
					Tax:               0.4,
					TaxRate:           0.1,
					Amount:            104.40,
					DueDate:           oneYearLater,
					Status:            "OUTSTANDING",
					CompanyID:         1,
					ServiceProviderID: 1,
				},
				{
					PaymentAmount:     100.00,
					Commission:        4.00,
					CommissionRate:    0.04,
					Tax:               0.4,
					TaxRate:           0.1,
					Amount:            104.40,
					DueDate:           oneYearLater,
					Status:            "OUTSTANDING",
					CompanyID:         1,
					ServiceProviderID: 2,
				},
			},
			error: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			service := NewService(nil, nil, tc.dbFactory(mockCtrl))

			got, err := service.ListInvoices(context.Background(), tc.req)
			if !tc.error {
				if err != nil {
					t.Errorf("want nil but error: %v", err.GetCode())
				}

				opts := cmpopts.IgnoreFields(model.Invoice{}, "IssueDate", "ID")
				if diff := cmp.Diff(got, tc.want, opts); diff != "" {
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
