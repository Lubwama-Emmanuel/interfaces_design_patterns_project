//nolint:(gocognit)
package filesystem_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
	filesystem "github.com/Lubwama-Emmanuel/Interfaces/storage/file_system"
)

func TestFileSytem(t *testing.T) { //nolint:(gocognit)
	t.Parallel()

	type args struct {
		number string
		data models.DataObject
	}

	file, err := ioutil.TempFile("", "temp.json")
	if err != nil {
		t.Fatalf("failed to create temp file %v", err.Error())
	}

	t.Cleanup(func() {
		defer os.Remove(file.Name())
	})

	tests := []struct {
		testName string
		filename string
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			filename: file.Name(),
			args: args{
				number: "0706039119",
				data: models.DataObject{
					"0706039119": "Lubwama",
				},
			},
			wantErr: assert.NoError,
		},
		{
			testName: "error",
			filename: "",
			args: args{
				number: "a",
				data: models.DataObject{
					"0706039119": "Lubwama",
				},
			},
			wantErr: assert.Error,
		},
		{
			testName: "error",
			filename: "",
			args:     args{},
			wantErr:  assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			fileDB := filesystem.NewFileSytemDatabase(tc.filename)

			err = fileDB.Create(tc.args.data)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, err))
				return
			}

			_, readErr := fileDB.Read(tc.args.number)
			if readErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readErr))
				return
			}

			updateErr := fileDB.Update(tc.args.data)
			if updateErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, updateErr))
				return
			}

			deleteErr := fileDB.Delete("1234567890")
			if deleteErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, deleteErr))
				return
			}

			_, readAllErr := fileDB.ReadAll()
			if readAllErr != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readAllErr))
				return
			}
		})
	}
}
