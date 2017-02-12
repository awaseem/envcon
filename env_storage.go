package main

import (
	"os"
)

// envStorage implements storer to interact with env
type envStorage struct{}

func (e *envStorage) set(key, value string) error {
	return os.Setenv(key, value)
}

func (e *envStorage) get(key string) (string, error) {
	return os.Getenv(key), nil
}
