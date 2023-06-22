package memory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/memory"
)

type args struct {
	number string
	data   models.DataObject
}

func TestMemory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "Success",
			args: args{
				number: "0706039119",
				data: models.DataObject{
					"0706039119": "Emmanuel",
				},
			},
			wantErr: assert.NoError,
		},
		{
			testName: "Error",
			args: args{
				number: "2",
				data:   models.DataObject{},
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			memoryDB := memory.NewMemoryStorage()

			performMemoryTest(t, tc, memoryDB)
		})
	}
}

func performMemoryTest(t *testing.T, tc struct {
	testName string
	args     args
	wantErr  assert.ErrorAssertionFunc
}, memoryDB *memory.MemoryDatabase,
) {
	createErr := memoryDB.Create(tc.args.data)
	if createErr != nil && tc.wantErr == nil {
		Helper(t, tc.testName, createErr)
		return
	}

	data, readErr := memoryDB.Read(tc.args.number)
	if readErr != nil && tc.wantErr == nil {
		Helper(t, tc.testName, readErr)
		return
	}

	if !assert.Equal(t, data, tc.args.data, "got data should be equal to expected data") {
		assert.Fail(t, fmt.Sprintf("Test %v data received %v yet expected %v", tc.testName, tc.args.data, data))
	}

	updateData := models.DataObject{
		"0706039119": "Lubwama",
	}

	updateErr := memoryDB.Update(updateData)
	if updateErr != nil && tc.wantErr == nil {
		Helper(t, tc.testName, updateErr)
		return
	}

	deleteErr := memoryDB.Delete(tc.args.number)
	if deleteErr != nil && tc.wantErr == nil {
		Helper(t, tc.testName, deleteErr)
		return
	}

	_, readAllErr := memoryDB.ReadAll()
	if readAllErr != nil && tc.wantErr == nil {
		Helper(t, tc.testName, readAllErr)
		return
	}
}

func Helper(t *testing.T, testName string, err error) {
	assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", testName, err))
}
