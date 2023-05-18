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
		path string
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
				path: "a",
				data: models.DataObject{
					"0706039119": "Emmanuel",
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			memoryDb := memory.NewMemoryStorage()
			createErr := memoryDb.Create(tc.args.path, tc.args.data)
			if createErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, createErr))
				return
			}

			memoryDb.Read(tc.args.path)
		})
	}

}
