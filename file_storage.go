package main

import "os"
import "io/ioutil"
import "encoding/json"

// fileStorage interact with json files for envs
type fileStorage struct {
	storageFolder string
}

// envFile implements storer to store and get env variables on disk
type envFile struct {
	file *os.File
	env  map[string]string
}

func (f *fileStorage) newFile(fileName string) (*envFile, error) {
	newF, err := os.Create(f.storageFolder + "/" + fileName)
	if err != nil {
		return nil, err
	}
	return &envFile{
		file: newF,
		env:  make(map[string]string),
	}, nil
}

func (f *fileStorage) getFile(fileName string) (*envFile, error) {
	openF, err := os.Open(f.storageFolder + "/" + fileName)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(openF)
	if err != nil {
		return nil, err
	}
	var envMap = make(map[string]string)
	if err = json.Unmarshal(b, &envMap); err != nil {
		return nil, err
	}
	return &envFile{
		file: openF,
		env:  envMap,
	}, nil
}

func (f *fileStorage) deleteFile(fileName string) error {
	return os.Remove(f.storageFolder + "/" + fileName)
}

func (f *fileStorage) listFiles() ([]string, error) {
	var fileNames []string
	fileInfos, err := ioutil.ReadDir(f.storageFolder)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		fileNames = append(fileNames, fileInfo.Name())
	}
	return fileNames, nil
}

func (e *envFile) set(key, value string) error {
	e.env[key] = value
	return nil
}

func (e *envFile) get(key string) (string, error) {
	return e.env[key], nil
}

func (e *envFile) save() error {
	b, err := json.Marshal(e.env)
	if err != nil {
		return err
	}
	_, err = e.file.Write(b)
	if err != nil {
		return err
	}
	return e.file.Sync()
}

func (e *envFile) close() error {
	return e.file.Close()
}
