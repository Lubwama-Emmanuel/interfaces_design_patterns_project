package postgres_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
	post "github.com/Lubwama-Emmanuel/Interfaces/storage/postgres"
)

type args struct {
	data models.DataObject
}

func TestPostgres(t *testing.T) {
	t.Parallel()

	mockDB, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})

	pgdb, err := post.NewPostgresDB(post.Config{}, dialector)
	if err != nil {
		t.Fatalf("failed to create postgres storage: %s", err)
	}

	storage := post.NewPhoneNumberStorage(pgdb)

	tests := []struct {
		testName string
		storage  *post.PhoneNumberStorage
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			storage:  storage,
			args: args{
				data: models.DataObject{
					"0704660968": "Emmanuel",
				},
			},
			wantErr: assert.NoError,
		},
		{
			testName: "success",
			storage:  storage,
			args: args{
				data: models.DataObject{
					"0706039119": "Lubwama",
				},
			},
			wantErr: assert.NoError,
		},
		{
			testName: "err database",
			storage:  storage,
			args: args{
				data: models.DataObject{
					"": "",
				},
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			performPostgresTest(t, tc, tc.storage)
		})
	}
}

func performPostgresTest(t *testing.T, tc struct {
	testName string
	storage  *post.PhoneNumberStorage
	args     args
	wantErr  assert.ErrorAssertionFunc
}, postgresDB *post.PhoneNumberStorage,
) {
	createErr := postgresDB.Create(tc.args.data)
	if createErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, createErr))
		return
	}

	_, readErr := postgresDB.Read("0704660968")
	if readErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readErr))
		return
	}

	updateObj := models.DataObject{
		"0704660968": "Rex",
	}

	updateErr := postgresDB.Update(updateObj)
	if updateErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, updateErr))
		return
	}

	deleteErr := postgresDB.Delete("0704660968")
	if deleteErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, deleteErr))
		return
	}

	_, readAllErr := postgresDB.ReadAll()
	if readAllErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readAllErr))
		return
	}
}
