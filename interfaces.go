package main

// generic interface for get method
type getter interface {
	get(key string) (string, error)
}

// generic interface for set method
type setter interface {
	set(key, value string) error
}
