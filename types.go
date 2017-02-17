package main

type fileContent struct {
	Encrypted bool   `json:"encrypted"`
	Salt      []byte `json:"salt"`
	Content   []byte `json:"content"`
}
