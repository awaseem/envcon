package main

import "os"
import "io/ioutil"
import "encoding/json"

// fileStorage interact with json files for envs
type fileStorage struct {
	concel        conceler
	storageFolder string
}

// envFile implements storer to store and get env variables on disk
type envFile struct {
	file        *os.File
	concel      conceler
	fileContent *fileContent
}

func (f *fileStorage) newFile(fileName string, encrypt bool) (*envFile, error) {
	newF, err := os.Create(f.storageFolder + "/" + fileName)
	if err != nil {
		return nil, err
	}
	return &envFile{
		file:   newF,
		concel: f.concel,
		fileContent: &fileContent{
			Encrypted: encrypt,
		},
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
	contents := fileContent{}
	if err = json.Unmarshal(b, &contents); err != nil {
		return nil, err
	}
	return &envFile{
		file:        openF,
		concel:      f.concel,
		fileContent: &contents,
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

func (e *envFile) setContent(env map[string]string, pass string) error {
	b, err := json.Marshal(env)
	if err != nil {
		return err
	}
	if e.fileContent.Encrypted {
		key, salt, err := e.concel.keyGen([]byte(pass))
		if err != nil {
			return err
		}
		encrypContent, err := e.concel.encrypt(key, b)
		if err != nil {
			return err
		}
		e.fileContent.Content = []byte(encrypContent)
		e.fileContent.Salt = salt
	} else {
		e.fileContent.Content = b
	}
	return nil
}

func (e *envFile) getContent(pass string) (map[string]string, error) {
	var env string
	var err error
	envMap := make(map[string]string)
	if e.fileContent.Encrypted {
		key := e.concel.keyGenWithSalt([]byte(pass), e.fileContent.Salt)
		env, err = e.concel.decrypt(key, string(e.fileContent.Content))
		if err != nil {
			return nil, err
		}
	} else {
		env = string(e.fileContent.Content)
	}
	if err := json.Unmarshal([]byte(env), &envMap); err != nil {
		return nil, err
	}
	return envMap, nil
}

func (e *envFile) save() error {
	b, err := json.Marshal(e.fileContent)
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
