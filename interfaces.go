package main

// generic interface for get/set method
type storer interface {
	get(key string) (string, error)
	set(key, value string) error
}
