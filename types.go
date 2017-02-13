package main

type settings struct {
	encrypt bool
	key     []byte
}

type fileContent struct {
	encrypted bool
	salt      string
	content   string
}
