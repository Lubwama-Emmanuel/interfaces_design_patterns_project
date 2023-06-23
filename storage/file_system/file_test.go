package filesystem_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Interfaces/models"
	filesystem "github.com/Lubwama-Emmanuel/Interfaces/storage/file_system"
)

type args struct {
	number string
	data   models.DataObject
}

func TestFileSytem(t *testing.T) {
	t.Parallel()

	f, err := os.Create("test.json")
	if err != nil {
		t.Fatalf("failed to create temp file %v", err.Error())
	}

	fileName := filepath.Base(f.Name())

	t.Cleanup(func() {
		defer os.Remove(fileName)
	})

	tests := []struct {
		testName string
		filename string
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			filename: fileName,
			args: args{
				number: "0706039119",
				data: models.DataObject{
					"0706039119": "Lubwama",
				},
			},
			wantErr: assert.NoError,
		},
		{
			testName: "error/wrong number",
			filename: "",
			args: args{
				number: "a",
				data:   models.DataObject{},
			},
			wantErr: assert.Error,
		},
		{
			testName: "error",
			filename: "",
			args: args{
				data: models.DataObject{},
			},
			wantErr: assert.Error,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			fileDB := filesystem.NewFileSytemDatabase(tc.filename)
			performFileTest(t, tc, fileDB)
		})
	}
}

func performFileTest(t *testing.T, tc struct {
	testName string
	filename string
	args     args
	wantErr  assert.ErrorAssertionFunc
}, fileDB *filesystem.FileSystemDatabase,
) {
	createErr := fileDB.Create(tc.args.data)
	if createErr != nil && tc.wantErr == nil {
		helper(t, tc.testName, createErr)
		return
	}

	_, readErr := fileDB.Read(tc.args.number)
	if readErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readErr))
		return
	}

	// if !assert.Equal(t, data, tc.args.data, "got data should be equal to expected data") {
	// 	fmt.Println("here", data, tc.args.data)
	// 	assert.Fail(t, fmt.Sprintf("Test %v data received %v yet expected %v", tc.testName, tc.args.data, data))
	// }

	updateErr := fileDB.Update(tc.args.data)
	if updateErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, updateErr))
		return
	}

	deleteErr := fileDB.Delete("0706039119")
	if deleteErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, deleteErr))
		return
	}

	_, readAllErr := fileDB.ReadAll()
	if readAllErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readAllErr))
		return
	}
}

func helper(t *testing.T, testName string, err error) bool {
	return assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", testName, err))
}
