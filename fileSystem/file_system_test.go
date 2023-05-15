package file_system_test

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemDatabase(t *testing.T) {
	file, err := ioutil.TempFile("", "database_test")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err.Error())
	}
	defer os.Remove(file.Name())

}
