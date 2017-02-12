package main

// generic interface for get/set method
type storer interface {
	get(key string) (string, error)
	set(key, value string) error
}

// generic interface for encrypting and decrypting
type conceler interface {
	encrypt(key, text []byte) (string, error)
	decrypt(key []byte, cryptoText string) (string, error)
}
