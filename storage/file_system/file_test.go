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
		data   models.DataObject
	}

	file, err := ioutil.TempFile("", "temp.json")
	// fileName := filepath.Base(file.Name())
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

			fileDB := filesystem.NewFileSytemDatabase("test.json")

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

// //nolint:(gocognit)
// package filesystem_test

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/Lubwama-Emmanuel/Interfaces/models"
// 	filesystem "github.com/Lubwama-Emmanuel/Interfaces/storage/file_system"
// )

// type args struct {
// 	number string
// 	data   models.DataObject
// }

// func TestFileSytem(t *testing.T) {
// 	t.Parallel()

// 	f, err := os.Create("test.json")
// 	// f, err := os.CreateTemp("", "tmp.json")
// 	// f, err := ioutil.TempFile("", "temp.json")
// 	if err != nil {
// 		t.Fatalf("failed to create temp file %v", err.Error())
// 	}

// 	// fileName := filepath.Base(f.Name())

// 	t.Cleanup(func() {
// 		// defer os.Remove("test.json")
// 	})

// 	tests := []struct {
// 		testName string
// 		filename string
// 		args     args
// 		wantErr  assert.ErrorAssertionFunc
// 	}{
// 		{
// 			testName: "success",
// 			filename: f.Name(),
// 			args: args{
// 				number: "0706039119",
// 				data: models.DataObject{
// 					"0706039119": "Lubwama",
// 				},
// 			},
// 			wantErr: assert.NoError,
// 		},
// 		{
// 			testName: "error/wrong number",
// 			filename: "",
// 			args: args{
// 				number: "a",
// 				data:   models.DataObject{},
// 			},
// 			wantErr: assert.Error,
// 		},
// 		{
// 			testName: "error",
// 			filename: "",
// 			args: args{
// 				data: models.DataObject{},
// 			},
// 			wantErr: assert.Error,
// 		},
// 	}

// 	for _, tc := range tests {
// 		tc := tc
// 		t.Run(tc.testName, func(t *testing.T) {
// 			t.Parallel()

// 			fileDB := filesystem.NewFileSytemDatabase("test.json")
// 			performFileTest(t, tc, fileDB)
// 		})
// 	}
// }

// func performFileTest(t *testing.T, tc struct {
// 	testName string
// 	filename string
// 	args     args
// 	wantErr  assert.ErrorAssertionFunc
// }, fileDB *filesystem.FileSystemDatabase,
// ) {
// 	createErr := fileDB.Create(tc.args.data)
// 	if createErr != nil && tc.wantErr == nil {
// 		helper(t, tc.testName, createErr)
// 		return
// 	}

// 	_, readErr := fileDB.Read(tc.args.number)
// 	if readErr != nil && tc.wantErr == nil {
// 		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readErr))
// 		return
// 	}

// 	// if !assert.Equal(t, data, tc.args.data, "got data should be equal to expected data") {
// 	// 	fmt.Println("here", data, tc.args.data)
// 	// 	assert.Fail(t, fmt.Sprintf("Test %v data received %v yet expected %v", tc.testName, tc.args.data, data))
// 	// }

// 	updateErr := fileDB.Update(tc.args.data)
// 	if updateErr != nil && tc.wantErr == nil {
// 		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, updateErr))
// 		return
// 	}

// 	deleteErr := fileDB.Delete("0706039119")
// 	if deleteErr != nil && tc.wantErr == nil {
// 		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, deleteErr))
// 		return
// 	}

// 	_, readAllErr := fileDB.ReadAll()
// 	if readAllErr != nil && tc.wantErr == nil {
// 		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readAllErr))
// 		return
// 	}
// }

// func helper(t *testing.T, testName string, err error) bool {
// 	return assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", testName, err))
// }
