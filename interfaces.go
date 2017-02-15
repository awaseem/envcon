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

type filer interface {
	newFile(fileName string, encrypt bool) (*envFile, error)
	getFile(fileName string) (*envFile, error)
	deleteFile(fileName string) error
	listFiles() ([]string, error)
}

type launcher interface {
	launch(envs map[string]string)
}

type commander interface {
	list() ([]string, error)
	source(fileName string) error
	create(fileName string, env map[string]string, encrypted bool) error
	update(fileName string, env map[string]string) error
	delete(fileName string) error
}
