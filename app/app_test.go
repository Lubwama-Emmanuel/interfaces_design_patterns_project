package app_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Interfaces/app"
	"github.com/Lubwama-Emmanuel/Interfaces/app/mocks"
	"github.com/Lubwama-Emmanuel/Interfaces/models"
)

func TestApp(t *testing.T) {
	t.Parallel()

	type args struct {
		name  string
		phone string
	}

	type fields struct {
		storage *mocks.MockIDatabase
	}

	tests := []struct {
		testName string
		prepare  func(t *testing.T, f *fields)
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			prepare: func(t *testing.T, f *fields) {
				f.storage.EXPECT().Create(models.DataObject{"0704660968": "Emmanuel"}).Return(nil)

				f.storage.EXPECT().Read("0704660968").Return(models.DataObject{"0704660968": "Emmanuel"}, nil)

				f.storage.EXPECT().Update(models.DataObject{"Emmanuel": "0704660968"}).Return(nil).AnyTimes()

				f.storage.EXPECT().Delete("0704660968").Return(nil).AnyTimes()

				f.storage.EXPECT().ReadAll().Return([]models.DataObject{
					{"0704660968": "Emmanuel"},
				}, nil)
			},
			args: args{
				name:  "Emmanuel",
				phone: "0704660968",
			},
			wantErr: assert.NoError,
		},
		{
			testName: "error",
			prepare: func(t *testing.T, f *fields) {
				f.storage.EXPECT().Create(gomock.Any()).Return(assert.AnError)

				f.storage.EXPECT().Read(gomock.Any()).Return(models.DataObject{}, assert.AnError)

				f.storage.EXPECT().Update(gomock.Any()).Return(assert.AnError)

				f.storage.EXPECT().Delete(gomock.Any()).Return(assert.AnError)

				f.storage.EXPECT().ReadAll().Return([]models.DataObject{}, assert.AnError)
			},
			args: args{
				name:  "Emmanuel",
				phone: "0704660968",
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			f := fields{
				storage: mocks.NewMockIDatabase(ctrl),
			}

			if tc.prepare != nil {
				tc.prepare(t, &f)
			}

			appInstance := app.NewApp(f.storage)

			err := appInstance.SavePhoneNumber(tc.args.name, tc.args.phone)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			_, err = appInstance.GetName(tc.args.phone)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			err = appInstance.UpdateName(tc.args.name, tc.args.phone)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			err = appInstance.DeleteContact(tc.args.phone)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			_, err =appInstance.GetAllPhoneNumbers()
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}
		})
	}
}
