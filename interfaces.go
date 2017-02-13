package main

// generic interface for get/set method
type storer interface {
	get(key string) (string, error)
	set(key, value string) error
}

// generic interface for encrypting and decrypting
type conceler interface {
	keyGen(pass []byte) (key []byte, salt []byte, err error)
	keyGenWithSalt(pass []byte, salt []byte) (key []byte)
	encrypt(key, text []byte) (string, error)
	decrypt(key []byte, cryptoText string) (string, error)
}
