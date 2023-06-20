package memory_test

import (
	"fmt"
	"testing"

	"github.com/Lubwama-Emmannuel/Interfaces/models"
	"github.com/Lubwama-Emmannuel/Interfaces/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	t.Parallel()

	type args struct {
		number string
		data models.DataObject
	}

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
				data: models.DataObject{},
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			memoryDB := memory.NewMemoryStorage()

			createErr := memoryDB.Create(tc.args.data)
			if createErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, createErr))
				return
			}

			data, readErr := memoryDB.Read(tc.args.number)
			if readErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readErr))
				return
			}

			if !compareMaps(data, tc.args.data) {
				assert.Fail(t, fmt.Sprintf("Test %v data received %v yet expected %v", tc.testName, tc.args.data, data))
			}

			updateData := models.DataObject{
				"0706039119": "Lubwama",
			}
			updateErr := memoryDB.Update(updateData)
			if updateErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, updateErr))
				return
			}

			deleteErr := memoryDB.Delete(tc.args.number)
			if deleteErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, deleteErr))
				return
			}

			_, readAllErr := memoryDB.ReadAll()
			if readAllErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readAllErr))
				return
			}
		})
	}
}

func compareMaps(map1, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; ok {
			if value1 != value2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
