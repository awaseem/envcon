package main

type settings struct {
	encrypt bool
	key     []byte
}

type fileContent struct {
	Encrypted bool   `json:"encrypted"`
	Salt      []byte `json:"salt"`
	Content   []byte `json:"content"`
}
