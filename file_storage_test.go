package main

import "testing"
import "os"

func TestCreateStore(t *testing.T) {
	// setup
	a := &aesCrypMock{}
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get work directory")
	}
	fs := fileStorage{
		concel:        a,
		storageFolder: wd + "/test",
	}

	// create directory
	err = fs.createStore()
	if err != nil {
		t.Error("createStore threw an error when it was not suppose too")
	}
	// check if directory was created
	if _, err := os.Stat(fs.storageFolder); os.IsNotExist(err) {
		t.Error("createStore did not create that directory")
	}

	// do not create directoy if it still exsits
	err = fs.createStore()
	if err != nil {
		t.Error("createStore threw and error when trying to check for directory")
	}

	// teardown
	err = os.Remove(fs.storageFolder)
	if err != nil {
		t.Errorf("failed to remove test directory")
	}
}

func TestExists(t *testing.T) {
	// setup
	a := &aesCrypMock{}
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get work directory")
	}
	fs := fileStorage{
		concel:        a,
		storageFolder: wd,
	}
	var tests = []struct {
		in  string
		out bool
	}{
		{in: "test.json", out: false},
		{in: "main.go", out: true},
	}
	for _, tt := range tests {
		e := fs.exists(tt.in)
		if e != tt.out {
			t.Error("exists returned unexpected value: (actual, expected)", e, tt.out)
		}
	}
}

func TestNewFile(t *testing.T) {
	// setup
	a := &aesCrypMock{}
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get work directory")
	}
	fs := fileStorage{
		concel:        a,
		storageFolder: wd,
	}
	var tests = []struct {
		fileName string
		encrypt  bool
		err      bool
	}{
		{"test.json", false, false},
		{"test.json", true, false},
	}
	for _, tt := range tests {
		envFile, err := fs.newFile(tt.fileName, tt.encrypt)
		if tt.err {
			if err == nil {
				t.Error("newFile did not throw an error when it was suppose too")
			}
		} else {
			if err != nil {
				t.Error("newFile threw en error when it was not suppose to")
			}
			if tt.encrypt != envFile.fileContent.Encrypted {
				t.Error("newFile encrypt flag does not match")
			}
		}
	}

	if err := os.Remove(wd + "/test.json"); err != nil {
		t.Errorf("failed to remove test file")
	}
}

func TestGetFile(t *testing.T) {
	// setup
	a := &aesCrypMock{}
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get work directory")
	}
	fs := fileStorage{
		concel:        a,
		storageFolder: wd + "/testData",
	}
	var tests = []struct {
		fileName string
		encrypt  bool
		err      bool
	}{
		{"test_no_exist.json", false, true},
		{"test.json", false, false},
		{"test_encrypt.json", true, false},
	}
	for _, tt := range tests {
		envFile, err := fs.getFile(tt.fileName)
		if tt.err {
			if err == nil {
				t.Error("getFile did not throw an error when it was suppose too")
			}
		} else {
			if err != nil {
				t.Error("getFile threw en error when it was not suppose to")
			}
			if tt.encrypt != envFile.fileContent.Encrypted {
				t.Error("getFile encrypt flag does not match")
			}
		}
	}
}

func TestDeleteFile(t *testing.T) {

}
