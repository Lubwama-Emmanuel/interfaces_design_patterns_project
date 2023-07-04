package postgres_test

import (
	"fmt"
	"testing"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/postgres"
	"github.com/stretchr/testify/assert"
)

type args struct {
	data models.DataObject
}

func TestPostgres(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		dbName   string
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			dbName:   "testdb",
			args: args{
				data: models.DataObject{
					"0704660968": "Emmanuel",
				},
			},
			wantErr: assert.NoError,
		},
		{
			testName: "success",
			dbName:   "",
			args: args{
				data: models.DataObject{
					"0706039119": "Rex",
				},
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			postgresDB := postgres.NewPostgresDB(tc.dbName)

			performPostgresTest(t, tc, postgresDB)
		})
	}
}

func performPostgresTest(t *testing.T, tc struct {
	testName string
	dbName   string
	args     args
	wantErr  assert.ErrorAssertionFunc
}, postgresDB *postgres.PostgresDB,
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
